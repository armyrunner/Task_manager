import {
  CTable,
  CFormSelect,
  CButton,
  CCard,
  CCardBody,
  CCardHeader,
  CFormLabel,
  CInputGroup,
  CModal,
  CModalHeader,
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
  const [showModal, setShowModal] = useState<boolean>(false);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [categoryLoading, setCategoryLoading] = useState<boolean>(false);
  const [searchQuery, setSearchQuery] = useState("");
  const [searchResults, setSearchResults] = useState<Task[]>([]);
  const [showSearchModal, setShowSearchModal] = useState<boolean>(false);
  const [reportType, setReportType] = useState<string>("");
  const [task, setTask] = useState<Task>({
    id: 0,
    description: "",
    due_date: "",
    start_date: "",
    finish_date: "",
    status: "pending",
    category_id: 0,
    category_name: "",
    notes: "",
  });

  const columns = [
    {
      key: "description",
      label: "Description",
      _props: { scope: "col" },
    },
  ];

  const items = [{}];

  //  =======  Fetch categories on component mount ==========
  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = async () => {
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
    }
  };

  // ============= HANDLE OPEN AND CLOSE MODALS

  const handleOpenModal = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setCategories("");
  };

    // ============= HANDLE CHANGE
  
    const handleChange = (
      e: React.ChangeEvent<
        HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
      >
    ) => {
      const { name, value } = e.target;
      // Convert category_id to number
      if (name === "category_id") {
        setTask((prev) => ({ ...prev, [name]: parseInt(value) || 0 }));
      } else {
        setTask((prev) => ({ ...prev, [name]: value }));
      }
    };
  return (
    <div className="d-flex justify-content-center align-items-start">
      <CCard>
        <CCardHeader>
          <CFormLabel>
            <strong>Select A Report</strong>
          </CFormLabel>
          <CFormSelect 
            id="reports" 
            name="reports"
            value={reportType}
            onChange={(e) => { 
                setReportType(e.target.value);
                if(e.target.value === "Category" ){
                    setShowModal(true);
                }
            }}
            >
            <option>Intial Task</option>
            <option>Category</option>
            <option>Completed Task</option>
            <option>Full Report</option>
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
          value={task.category_id}
          onChange={handleChange}
        >
          <option value={0}>Select Category</option>
          {categories.map((cat) => (
            <option key={cat.id} value={cat.id}>
              {cat.name}
            </option>
          ))}
        </CFormSelect>
        <CButton color="success">
            <CIcon icon={cilThumbUp}/>
            OK
        </CButton>
      </CModal>
    </div>
  );
}

export default Reports;
