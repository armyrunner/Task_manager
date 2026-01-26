import {
  CSidebar,
  CSidebarHeader,
  CSidebarNav,
  CSidebarToggler,
  CNavItem,
  CNavGroup,
  CButton,
  CInputGroup,
  CFormInput,
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
  cilSearch,
  cilHome,
} from "@coreui/icons";
import "@coreui/coreui/dist/css/coreui.min.css";
import styles from "./TaskDashboard.module.css";
import { useState, useEffect } from "react";
import { Outlet, Link } from "react-router-dom";

const Dashboard = () => {
  const [categories, setCategories] = useState<string[]>([]);

  useEffect(() => {
    // Fetch categories when component mounts
    fetch("http://localhost:8080/api/categories",{
      method: "GET",
      headers:{
        'Content-Type': 'application/json',
        Authorization: `Bearer ${localStorage.getItem('access_token')}`
      }
    })
      .then((res) => res.json())
      .then((data) => setCategories(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div className={styles.pageWrapper}>
      <CSidebar className="border-end">
        <CSidebarHeader className="border-bottom">
          <CInputGroup>
            <CFormInput
              type="search"
              placeholder="Search ID or Description"
              aria-label="Search"
            />
            <CButton type="button" color="primary" variant="outline">
              <CIcon icon={cilSearch} />
            </CButton>
          </CInputGroup>
        </CSidebarHeader>
        <CSidebarNav>
          <CNavItem>
            <Link to="/taskdashboard" className="nav-link">
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
              <CNavItem key={category.id} href="#">
                <span className="nav-icon">
                  <span className="nav-icon-bullet"></span>
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
