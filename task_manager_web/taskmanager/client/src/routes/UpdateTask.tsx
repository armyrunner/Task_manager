import {
  CForm,
  CFormInput,
  CButton,
  CFormLabel,
  CFormSelect,
  CFormTextarea,
  CCol,
  CRow,
  CCard,
  CCardHeader,
  CCardBody,
  CCardFooter,
  CModal,
  CModalHeader,
  CModalBody,
  CModalTitle,
  CModalFooter,
  CInputGroup,
  CTooltip,
  CSpinner,
} from "@coreui/react";
import { cilSave, cilX, cilSearch, cilPlus } from "@coreui/icons";
import CIcon from "@coreui/icons-react";
import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

interface Category {
  id: number;
  user_id: number;
  name: string;
}

function UpdateTask() {
  const navigate = useNavigate();
  const [showModal, setShowModal] = useState(false);
  const [newCategory, setNewCategory] = useState("");
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [categoryLoading, setCategoryLoading] = useState<boolean>(false);
  const [searchQuery, setSearchQuery] = useState("");

  const [task, setTask] = useState({
    id: 0,
    description: "",
    due_date: "",
    start_date: "",
    finish_date: "",
    status: "pending",
    category_id: 0,
    notes: "",
  });

  // Fetch categories on component mount
  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/categories", {
        method: "GET",
        headers: {
          "Content-Type": 'application/json',
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

  const handleOpenModal = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setNewCategory("");
  };

  const handleSaveCategory = async () => {
    if (newCategory.trim() === "") return;

    setCategoryLoading(true);
    try {
      const response = await fetch("http://localhost:8080/api/categories", {
        method: "POST",
        headers: {
          "Content-Type": 'application/json',
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
        body: JSON.stringify({ name: newCategory }),
      });

      if (response.ok) {
        const newCat = await response.json();
        setCategories([...categories, newCat]);
        setTask({ ...task, category_id: newCat.id });
        handleCloseModal();
      } else {
        const errorData = await response.json();
        setError(errorData.message || "Failed to create category");
      }
    } catch (err) {
      console.error("Error creating category:", err);
      setError("Failed to create category");
    } finally {
      setCategoryLoading(false);
    }
  };

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

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("Updated Task:", task);

    if (task.id === 0) {
      setError("Please search for a task first");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const response = await fetch("http://localhost:8080/api/tasks", {
        method: "PUT",
        headers: {
          "Content-Type": 'application/json',
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
        body: JSON.stringify(task)
      });
      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.message || data.Error || "Failed to update task");
      }
      navigate("/taskdashboard", { replace: true });
    } catch (error) {
      console.error("Error updating task:", error);
      setError("Failed to update task. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  const handleCancel = () => {
    navigate("/taskdashboard");
  };

  const handleSearch = async () => {
    if (!searchQuery.trim()) return;

    setLoading(true);

    try{
      const resp = await fetch(
        `http://localhost:8080/api/tasks?search=${encodeURIComponent(searchQuery)}`,{
          headers:{
            'Content-Type':'application/json',
            Authorization: `Bearer ${localStorage.getItem("access_token")}`
          },
        }
      );
      const data = await resp.json();
      console.log("Search response",data)
      let taskFound = null;

      if (Array.isArray(data) && data.length > 0){
        taskFound = data;
      } else if (data.tasks && Array.isArray(data.tasks) && data.tasks.length > 0){
        taskFound = data.tasks;
      }

      if(taskFound){
        setTask(taskFound)
      } else {
        setError("Task not found!")
      }
    } catch (err) {
      setError("Search failed");
      console.error("Search failed:", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="d-flex justify-content-center align-items-start w-100 h-100 pt-4">
      <CCard style={{ maxWidth: "600px", width: "100%" }}>
        <CCardHeader className="d-flex justify-content-between align-items-center">
          <strong>Update Task</strong>
          <div className="d-flex gap-2">
            <CFormInput
              type="search"
              value={searchQuery}
              placeholder="Search By Description"
              aria-label="Search"
              style={{ maxWidth: "200px" }}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            <CButton type="button" color="primary" variant="outline" onClick={handleSearch}>
              <CIcon icon={cilSearch} />
            </CButton>
          </div>
        </CCardHeader>

        <CForm onSubmit={handleSubmit}>
          {loading && <CSpinner color="success" />}
          {error && <div className="alert alert-danger">{error}</div>}
          <CCardBody>
            <CRow>
              <CCol md={12} className="mb-3">
                <CFormLabel htmlFor="description">Task Name</CFormLabel>
                <CFormInput
                  type="text"
                  id="description"
                  name="description"
                  value={task.description}
                  onChange={handleChange}
                  required
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="start_date">Start Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="start_date"
                  name="start_date"
                  value={task.start_date}
                  onChange={handleChange}
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="finish_date">Finish Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="finish_date"
                  name="finish_date"
                  value={task.finish_date}
                  onChange={handleChange}
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="category_id">Category</CFormLabel>
                <CInputGroup>
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
                  <CTooltip content="Add New Category" placement="top">
                    <CButton
                      color="primary"
                      variant="outline"
                      onClick={handleOpenModal}
                    >
                      <CIcon icon={cilPlus} />
                    </CButton>
                  </CTooltip>
                </CInputGroup>
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="due_date">Due Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="due_date"
                  name="due_date"
                  value={task.due_date}
                  onChange={handleChange}
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="status">Status</CFormLabel>
                <CFormSelect
                  id="status"
                  name="status"
                  value={task.status}
                  onChange={handleChange}
                >
                  <option value="pending">Pending</option>
                  <option value="in-progress">In-Progress</option>
                  <option value="completed">Completed</option>
                  <option value="cancelled">Canceled</option>
                </CFormSelect>
              </CCol>
              <CCol md={12} className="mb-3">
                <CFormLabel htmlFor="notes">Notes</CFormLabel>
                <CFormTextarea
                  id="notes"
                  name="notes"
                  rows={4}
                  value={task.notes}
                  onChange={handleChange}
                  style={{ resize: "none" }}
                />
              </CCol>
            </CRow>
          </CCardBody>
          <CCardFooter className="d-flex justify-content-end gap-2">
            <CButton color="danger" variant="outline" onClick={handleCancel}>
              <CIcon icon={cilX} className="me-2" />
              Cancel
            </CButton>
            <CButton color="primary" type="submit" disabled={loading}>
              <CIcon icon={cilSave} className="me-2" />
              Save Task
            </CButton>
          </CCardFooter>
        </CForm>
      </CCard>
      <CModal visible={showModal} onClose={handleCloseModal}>
        <CModalHeader>
          <CModalTitle>Add New Category</CModalTitle>
        </CModalHeader>
        <CModalBody>
          <CFormInput
            type="text"
            id="newCategory"
            name="newCategory"
            value={newCategory}
            onChange={(e) => setNewCategory(e.target.value)}
            placeholder="Enter new category"
          />
        </CModalBody>
        <CModalFooter>
          <CButton
            color="secondary"
            variant="outline"
            onClick={handleCloseModal}
          >
            Cancel
          </CButton>
          <CButton
            color="primary"
            onClick={handleSaveCategory}
            disabled={categoryLoading}
          >
            {categoryLoading ? <CSpinner size="sm" /> : "Save"}
          </CButton>
        </CModalFooter>
      </CModal>
    </div>
  );
}

export default UpdateTask;
