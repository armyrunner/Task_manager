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
} from "@coreui/react";
import {
  CModal,
  CModalHeader,
  CModalBody,
  CModalTitle,
  CModalFooter,
  CInputGroup,
  CSpinner,
} from "@coreui/react";
import { cilSave, cilX,cilPlus, cilList } from "@coreui/icons";
import CIcon from "@coreui/icons-react";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { CTooltip } from "@coreui/react";

function AddTask() {
  const navigate = useNavigate();
  const [showModal, setShowModal] = useState(false);
  const [loading,setLoading] = useState<boolean>(false);
  const [error,setError] = useState<string | null>(null);
  const [newCategory, setNewCategory] = useState("");
  const [categories, setCategories] = useState<{value: string, label: string}[]>([
    {value: "personal", label: "Personal"},
    {value: "work", label: "Work"},
    {value: "family", label: "Family"},
    {value: "other", label: "Other"},
  ]);
  const [showSuccessModal, setShowSuccessModal] = useState<boolean>(false);

  
  const [task, setTask] = useState({
    description: "",
    due_date: "",
    start_date: "",
    finish_date: "",
    status: "pending",
    category: "",
    notes: "",
  });

  const handleOpenModal = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setNewCategory(""); // Clear the new category input
  };

  const handleSaveCategory = () => {
    if (newCategory.trim()=== "") return;

    const newCat = {
      value: newCategory.toLowerCase().replace(/\s+/g, '-'),
      label: newCategory,
    };
  
    setCategories([...categories, newCat]);  // Add to list
    setTask({ ...task, category: newCat.value });  // Select it
    handleCloseModal();
  };

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >
  ) => {
    const { name, value } = e.target;
    setTask((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("New Task:", task);

    setLoading(true);
    setError(null);

    if(!task.description){
      setError('Task description is required');
      setLoading(false);
      return;
    }
    
    // TODO: Send to API
    try{
      const response = await fetch('http://localhost:8080/api/tasks',{
        method: 'POST',
        headers: {'Content-Type': 'application/json','Authorization': `Bearer ${localStorage.getItem('access_token')}`},
        body: JSON.stringify(task),
        credentials: 'include'
      });
      const data = await response.json();
      if(!response.ok){
        throw new Error(data.message || data.Error || 'Failed to add task');
      }
      setTask({
        description: "",
        due_date: "",
        start_date: "",
        finish_date: "",
        status: "pending",
        category: "",
        notes: "",
      });
      setShowSuccessModal(true);
    } catch (error) {
      console.error('Error adding task:', error);
      setError('Failed to add task. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleCancel = () => {
    navigate("/taskdashboard");
  };

  const handleGoToDashboard = () => {
    setShowSuccessModal(false);
    navigate("/taskdashboard",{replace: true});
  };

  const handleAddAnotherTask = () => {
    setShowSuccessModal(false);
    setTask({
      description: "",
      due_date: "",
      start_date: "",
      finish_date: "",
      status: "pending",
      category: "",
      notes: "",  
    });
    setLoading(false);
    setError(null);
  };

  return (
    <div className="d-flex justify-content-center align-items-start w-100 h-100 pt-4">
      <CCard style={{ maxWidth: "600px", width: "100%" }}>
        <CCardHeader>
          <strong>Add New Task</strong>
        </CCardHeader>

        <CForm onSubmit={handleSubmit}>
          {loading && <CSpinner color="success" />}
          {error && <div className="alert alert-danger">{error}</div>}
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
                  required
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="startDate">Start Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="start_date"
                  name="start_date"
                  value={task.start_date}
                  onChange={handleChange}
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="finishDate">Finish Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="finish_date"
                  name="finish_date"
                  value={task.finish_date}
                  onChange={handleChange}
                />
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="category">Category</CFormLabel>
                <CInputGroup>
                <CFormSelect
                  id="category"
                  name="category"
                  value={task.category}
                  onChange={handleChange}
                >
                  {categories.map((cat) => (
                    <option key={cat.value} value={cat.value}>
                      {cat.label}
                    </option>
                  ))}
                </CFormSelect>
                <CTooltip content="Add New Category" placement="top">
                  <CButton color="primary" variant="outline" onClick={handleOpenModal}>
                    <CIcon icon={cilPlus} className="me-2" />
                  </CButton>
                </CTooltip>
                </CInputGroup>
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="dueDate">Due Date</CFormLabel>
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
            <CButton color="primary" type="submit">
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

          <CButton color="secondary" variant="outline" onClick={handleCloseModal}>
            Cancel
          </CButton>
          <CButton color="primary" onClick={handleSaveCategory}>
            Save
          </CButton>
        </CModalFooter>
      </CModal>
      <CModal visible={showSuccessModal}>
        <CModalHeader>
          <CModalTitle>Task Added Successfully</CModalTitle>
        </CModalHeader>
        <CModalBody>
          <p>Task added successfully. You can now add another task.</p>
        </CModalBody>
        <CModalFooter>
          <CButton color="primary" onClick={handleAddAnotherTask}>
            <CIcon icon={cilPlus} className="me-2" />
            Add Another Task
          </CButton>
          <CButton color="primary" onClick={handleGoToDashboard}>
            <CIcon icon={cilList} className="me-2" />
            Go to Dashboard
          </CButton>
        </CModalFooter>
      </CModal>
    </div>
  );
}

export default AddTask;
