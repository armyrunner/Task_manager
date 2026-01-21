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
  CModalTitle,
  CModalBody,
  CModalFooter,
  CSpinner,
} from "@coreui/react";
import { cilTrash, cilX, cilSearch, cilWarning } from "@coreui/icons";
import CIcon from "@coreui/icons-react";
import React, {useState } from "react";
import { useNavigate } from "react-router-dom";


function DeleteTask() {
  const navigate = useNavigate();
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [loading, setLoading] = useState<boolean>(false);
  const [searchQuery, setSearchQuery] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [task, setTask] = useState({
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

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >
  ) => {
    const { name, value } = e.target;
    setTask((prev) => ({ ...prev, [name]: value }));
  };

  const handleDeleteClick = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setShowConfirmModal(true);
  };

  const handleSearch = async () => {
    if (!searchQuery.trim()) return;

    setLoading(true);

    try{
      const response = await fetch(
        `http://localhost:8080/api/tasks?search=${encodeURIComponent(searchQuery)}`,{
          headers:{
            Authorization: `Bearer ${localStorage.getItem("access_token")}`
          },
        }
      );
      const data = await response.json();
      if (response.ok && data.length > 0){
        setTask(data[0]);
      } else {
        setError("Task not Found");
      }
    } catch (error) {
      setError("Search failed");
      console.error("Search failed:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleConfirmDelete = async () => {
    console.log("Delete Task:", task);
    
    try{
      const response = await fetch(`http://localhost:8080/api/tasks/${task.id}`,{
        method: "DELETE",
        headers:{
          'Authorization': `Bearer ${localStorage.getItem("access_token")}`
        },
        credentials:"include"
      });
      setTask({
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
      const data = await response.json();
      if(!response.ok){
        throw new Error(data.message || data.Error || 'Failed to delete task')
      }
      setShowConfirmModal(false);
      navigate("/taskdashboard");
    } catch(error){
      console.error("Failed to delete", error)
      setError("Failed to Delete Task!")
    } finally {
      setLoading(false);
    }
    
  };

  const handleCancelDelete = () => {
    setShowConfirmModal(false);
  };

  const handleCancel = () => {
    navigate("/taskdashboard");
  };

  return (
    <div className="d-flex justify-content-center align-items-start w-100 h-100 pt-4">
      {loading && <CSpinner  color="success" />}
      {error && <div className="alert alert-danger">{error}</div>}
      <CCard style={{ maxWidth: "600px", width: "100%" }}>
        <CCardHeader className="d-flex justify-content-between align-items-center">
          <strong>Delete Task</strong>
          <div className="d-flex gap-2">
            <CFormInput
              type="search"
              placeholder="Search..."
              aria-label="Search"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              style={{ maxWidth: "200px" }}
            />
            <CButton type="button" color="primary" variant="outline" onClick={handleSearch}>
              <CIcon icon={cilSearch} />
            </CButton>
          </div>
        </CCardHeader>

        <CForm onSubmit={handleDeleteClick}>
          <CCardBody>
            <CRow>
              <CCol md={12} className="mb-3">
                <CFormLabel htmlFor="name">Task Name</CFormLabel>
                <CFormInput
                  type="text"
                  id="description"
                  name="description"
                  value={task.description}
                  onChange={handleChange}
                  disabled
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
                  disabled
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
                  disabled
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="category">Category</CFormLabel>
                <CFormInput
                  type="text"
                  id="category"
                  value={task.category_name}
                  disabled />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="due_date">Due Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="due_date"
                  name="due_date"
                  value={task.due_date}
                  onChange={handleChange}
                  disabled
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="status">Status</CFormLabel>
                <CFormSelect
                  id="status"
                  name="status"
                  value={task.status}
                  onChange={handleChange}
                  disabled
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
                  disabled
                />
              </CCol>
            </CRow>
          </CCardBody>
          <CCardFooter className="d-flex justify-content-end gap-2">
            <CButton color="secondary" variant="outline" onClick={handleCancel}>
              <CIcon icon={cilX} className="me-2" />
              Cancel
            </CButton>
            <CButton color="danger" type="submit">
              <CIcon icon={cilTrash} className="me-2" />
              Delete Task
            </CButton>
          </CCardFooter>
        </CForm>
      </CCard>

      {/* Confirmation Modal */}
      <CModal visible={showConfirmModal} onClose={handleCancelDelete}>
        <CModalHeader>
          <CModalTitle>Confirm Delete</CModalTitle>
        </CModalHeader>
        <CModalBody className="text-center">
          <CIcon icon={cilWarning} size="3xl" className="text-danger mb-3" />
          <p className="mb-0">Are you sure you want to delete this task?</p>
          <p className="text-muted">
            <strong>{task.description || "This task"}</strong>
          </p>
          <p className="text-danger small">This action cannot be undone.</p>
        </CModalBody>
        <CModalFooter>
          <CButton
            color="secondary"
            variant="outline"
            onClick={handleCancelDelete}
          >
            Cancel
          </CButton>
          <CButton color="danger" onClick={handleConfirmDelete}>
            <CIcon icon={cilTrash} className="me-2" />
            Yes, Delete
          </CButton>
        </CModalFooter>
      </CModal>
    </div>
  );
}

export default DeleteTask;

