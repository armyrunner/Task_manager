import {
  CTable,
  CFormSelect,
  CButton,
  CCard,
  CCardBody,
  CCardHeader,
  CFormLabel,
  CModal,
  CModalHeader,
  CSpinner,
} from "@coreui/react";

import {cilPrint, cilThumbUp } from "@coreui/icons";
import CIcon from "@coreui/icons-react";
import { useEffect, useState } from "react";


interface Category {
    id: number;
    user_id: number;
    name: string;
  }
  
  interface Task {
    id: number;
    description: string;
    due_date: string;
    start_date: string;
    finish_date: string;
    status: string;
    category_id: number;
    category_name: string;
    notes: string;
  }


function Reports() {
  // Modal state
  const [showModal, setShowModal] = useState<boolean>(false);

  // Data state
  const [categories, setCategories] = useState<Category[]>([]);
  const [items, setItems] = useState<Task[]>([]);

  // Selection state
  const [reportType, setReportType] = useState<string>("");
  const [selectedCategoryId, setSelectedCategoryId] = useState<number>(0);

  // Loading/Error state
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const columnConfig: Record<string, string> = {
    description: "Description",
    due_date: "Due Date",
    start_date: "Start Date",
    finish_date: "Finish Date",
    status: "Status",
    category_name: "Category",
    notes: "Notes",

  }

  const columns = Object.entries(columnConfig).map(([key, label]) => ({
    key,
    label,
    _props: {scope: "col"},
  }))


  //  =======  Fetch categories on component mount ==========
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

  const fetchTasksByCategory = async (categoryID: number) => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(`http://localhost:8080/api/tasks?category_id=${categoryID}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
      });
      if (response.ok) {
        const data = await response.json();
        return data;
      }
    } catch (err) {
      console.error("Error fetching categories:", err);
    }
  };

  const fetchInitialTasks = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch("http://localhost:8080/api/tasks", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
      });
      if (response.ok) {
        const data = await response.json();
        return data;
      }
    } catch (err) {
      console.error("Error fetching initial tasks:", err);
    }
  };

  const fetchCompletedTasks = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch("http://localhost:8080/api/tasks/completed", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
      });
      if (response.ok) {
        const data = await response.json();
        return data;
      }
    } catch (err) {
      console.error("Error fetching initial tasks:", err);
    }
  };

  const fetchAllTasks = async () => {
    setLoading(true);
    setError(null);
    const initial = await fetchInitialTasks();
    const completed = await fetchCompletedTasks();
    return [...(initial || []), ...(completed || [])];
  };

  // ============= HANDLE OPEN AND CLOSE MODALS

  const handleOpenModal = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedCategoryId(0);
  };

    // ============= HANDLE CHANGE
  
    const handleReportChange = async (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selected = e.target.value;
        setReportType(selected);
        setItems([]); // Clear previous results

        if (selected === "Category") {
            handleOpenModal();
        } else if (selected === "Initial Tasks") {
            const data = await fetchInitialTasks();
            setItems(data || []);
            setLoading(false);
        } else if (selected === "Completed Tasks") {
            const data = await fetchCompletedTasks();
            setItems(data || []);
            setLoading(false);
        } else if (selected === "Full Report") {
            const data = await fetchAllTasks();
            setItems(data || []);
            setLoading(false);
        }
    };

    const handleCategorySelect = async () => {
        if (selectedCategoryId === 0) return;
        const data = await fetchTasksByCategory(selectedCategoryId);
        setItems(data);
        setShowModal(false);
        setSelectedCategoryId(0);
        setLoading(false);
    }
  return (
    <div className="d-flex justify-content-center align-items-start">
        {loading && <CSpinner color="success" />}
        {error && <div className="alert alert-danger">Error: {error}</div>}
      <CCard>
        <CCardHeader>
          <CFormLabel>
            <strong>Select A Report</strong>
          </CFormLabel>
          <CFormSelect 
            id="reports" 
            name="reports"
            value={reportType}
            onChange={handleReportChange}
            >
            <option value="">Select a Report</option>
            <option value="Initial Tasks">Initial Tasks</option>
            <option value="Category">Category</option>
            <option value="Completed Tasks">Completed Tasks</option>
            <option value="Full Report">Full Report</option>
          </CFormSelect>
        </CCardHeader>
        <CCardBody>
          <CTable
            columns={columns}
            items={items}
            tableHeadProps={{ color: "light" }}
            striped
            hover
            bordered
          />
          <CButton>
            <CIcon icon={cilPrint} className="me-2" />
            Print
          </CButton>
          <CButton
            color="warning"
            variant="outline"
            onClick={handleCloseModal}
          >
            Cancel
          </CButton>
        </CCardBody>
      </CCard>
      <CModal visible={showModal} onClose={handleCloseModal}>
        <CModalHeader>
          <strong>Choose a category</strong>
        </CModalHeader>
        <CFormSelect
          id="category_id"
          name="category_id"
          value={selectedCategoryId}
          onChange={(e) => setSelectedCategoryId(parseInt(e.target.value) || 0)}
          className="m-3"
        >
          <option value={0}>Select Category</option>
          {categories.map((cat) => (
            <option key={cat.id} value={cat.id}>
              {cat.name}
            </option>
          ))}
        </CFormSelect>
        <CButton color="success" onClick={handleCategorySelect} className="m-3">
            <CIcon icon={cilThumbUp} className="me-2"/>
            OK
        </CButton>
      </CModal>
    </div>
  );
}

export default Reports;
