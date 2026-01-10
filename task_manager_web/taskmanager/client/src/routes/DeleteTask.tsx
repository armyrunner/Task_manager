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
  } from "@coreui/react";
import { cilTrash, cilX, cilSearch, cilWarning } from "@coreui/icons";
  import CIcon from "@coreui/icons-react";
  import React, { useState } from "react";
  import { useNavigate } from "react-router-dom";
  
function DeleteTask() {
    const navigate = useNavigate();
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [categories] = useState<{value: string, label: string}[]>([
    {value: "personal", label: "Personal"},
    {value: "work", label: "Work"},
    {value: "family", label: "Family"},
    {value: "other", label: "Other"},
  ]);

    const [task, setTask] = useState({
      name: "",
      dueDate: "",
      startDate: "",
      finishDate: "",
      status: "pending",
    category: "",
      notes: "",
    });
  
    const handleChange = (
      e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>
    ) => {
      const { name, value } = e.target;
      setTask((prev) => ({ ...prev, [name]: value }));
    };
  
  const handleDeleteClick = (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
    setShowConfirmModal(true);
  };

  const handleConfirmDelete = () => {
    console.log("Delete Task:", task);
      // TODO: Send to API
      setTask({
        name: "",
        dueDate: "",
        startDate: "",
        finishDate: "",
        status: "pending",
      category: "",
        notes: "",
      });
    setShowConfirmModal(false);
    navigate("/taskdashboard");
  };

  const handleCancelDelete = () => {
    setShowConfirmModal(false);
    };
  
    const handleCancel = () => {
      navigate("/taskdashboard");
    };
  
    return (
      <div className="d-flex justify-content-center align-items-start w-100 h-100 pt-4">
      <CCard style={{ maxWidth: '600px', width: '100%' }}>
        <CCardHeader className="d-flex justify-content-between align-items-center">
          <strong>Delete Task</strong>
          <div className="d-flex gap-2">
            <CFormInput
              type="search"
              placeholder="Search..."
              aria-label="Search"
              style={{ maxWidth: '200px' }}
            />
            <CButton type="button" color="primary" variant="outline">
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
                  id="name"
                  name="name"
                  value={task.name}
                  onChange={handleChange}
                  disabled
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="startDate">Start Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="startDate"
                  name="startDate"
                  value={task.startDate}
                  onChange={handleChange}
                  disabled
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="finishDate">Finish Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="finishDate"
                  name="finishDate"
                  value={task.finishDate}
                  onChange={handleChange}
                  disabled
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="category">Category</CFormLabel>
                <CFormSelect
                  id="category"
                  name="category"
                  value={task.category}
                  onChange={handleChange}
                  disabled
                >
                  {categories.map((cat) => (
                    <option key={cat.value} value={cat.value}>
                      {cat.label}
                    </option>
                  ))}
                </CFormSelect>
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="dueDate">Due Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="dueDate"
                  name="dueDate"
                  value={task.dueDate}
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
                  style={{ resize: 'none' }}
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
          <p className="text-muted"><strong>{task.name || "This task"}</strong></p>
          <p className="text-danger small">This action cannot be undone.</p>
        </CModalBody>
        <CModalFooter>
          <CButton color="secondary" variant="outline" onClick={handleCancelDelete}>
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
  