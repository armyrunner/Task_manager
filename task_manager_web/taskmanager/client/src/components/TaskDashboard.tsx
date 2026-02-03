import {
  CSidebar,
  CSidebarHeader,
  CSidebarNav,
  CSidebarToggler,
  CNavItem,
  CNavGroup,
  CSpinner
} from "@coreui/react";

import CIcon from "@coreui/icons-react";
import {
  cilList,
  cilPlus,
  cilPenAlt,
  cilTrash,
  cilPuzzle,
  cilCheckCircle,
  cilSitemap,
  cilHome,
  cilSpreadsheet,
} from "@coreui/icons";
import "@coreui/coreui/dist/css/coreui.min.css";
import styles from "./TaskDashboard.module.css";
import { useState, useEffect } from "react";
import { Outlet, Link } from "react-router-dom";


const Dashboard = () => {
  const [categories, setCategories] = useState<string[]>([]);

  // Loading/Error state
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchCategories();    
  }, []);

  const fetchCategories = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch("http://localhost:8080/api/categories", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
      });
      if (response.ok) {
        const data = await response.json();
        setCategories(data || []);
      }
    } catch (err) {
      console.error("Error fetching categories:", err);
    } finally {
      setLoading(false);
    }
  };


  return (
    <div className={styles.pageWrapper}>
      {loading && <CSpinner color="success" />}
      {error && <div className="alert alert-danger">Error: {error}</div>}
      <CSidebar className="border-end">
        <CSidebarHeader className="border-bottom">
        </CSidebarHeader>
        <CSidebarNav>
          <CNavItem>
            <Link to="taskdashboard" className="nav-link">
              <CIcon customClassName="nav-icon text-primary" icon={cilHome} />{" "}
              Dashboard
            </Link>
          </CNavItem>
          <CNavItem>
            <Link to="addcategory" className="nav-link">
              <CIcon
                customClassName="nav-icon text-primary"
                icon={cilSitemap}
              />
              Add Category
            </Link>
          </CNavItem>
          <CNavItem>
            <Link to="addtask" className="nav-link">
              <CIcon customClassName="nav-icon text-primary" icon={cilPlus} />{" "}
              Add Task
            </Link>
          </CNavItem>
          <CNavItem>
            <Link to="updatetask" className="nav-link">
              <CIcon customClassName="nav-icon text-info" icon={cilPenAlt} />{" "}
              Update Task
            </Link>
          </CNavItem>
          <CNavItem>
            <Link to="deletetask" className="nav-link">
              <CIcon customClassName="nav-icon text-danger" icon={cilTrash} />{" "}
              Delete Task
            </Link>
          </CNavItem>
          <CNavItem href="#">
            <CIcon customClassName="nav-icon text-info" icon={cilPuzzle} />
            All Tasks
          </CNavItem>
          <CNavGroup
            toggler={
              <>
                <CIcon
                  customClassName="nav-icon text-secondary"
                  icon={cilList}
                />{" "}
                Categories
              </>
            }
          >
            {categories.map((category: any) => (
              <CNavItem key={category.id} to={`/taskdashboard?category_id=${category.id}&category_name=${category.name}`}>
                <span className="nav-icon">
                </span>{" "}
                {category.name}
              </CNavItem>
            ))}
          </CNavGroup>
          <CNavItem href="#">
            <CIcon
              customClassName="nav-icon text-success"
              icon={cilCheckCircle}
            />
            Completed Tasks
          </CNavItem>
          {/* Need to Add report drop down Initial / Category (possibly sub-category) / Completed Tasks */}
          <CNavItem>
            <Link to="reports" className="nav-link">
              <CIcon
                customClassName="nav-icon text-primary"
                icon={cilSpreadsheet}
              />{" "}
              Print Reports
            </Link>
          </CNavItem>
        </CSidebarNav>
        <CSidebarHeader className="border-top">
          <CSidebarToggler />
        </CSidebarHeader>
      </CSidebar>
      <div className={styles.mainContent}>
        <Outlet />
      </div>
    </div>
  );
};

export default Dashboard;
