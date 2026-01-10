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
  } from "@coreui/react";
import { cilSave, cilX, cilSearch, cilPlus } from "@coreui/icons";
  import CIcon from "@coreui/icons-react";
  import React, { useState } from "react";
  import { useNavigate } from "react-router-dom";
  
function UpdateTask() {
    const navigate = useNavigate();
  const [showModal, setShowModal] = useState(false);
  const [newCategory, setNewCategory] = useState("");
  const [categories, setCategories] = useState<{value: string, label: string}[]>([
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

  const handleOpenModal = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setNewCategory("");
  };

  const handleSaveCategory = () => {
    if (newCategory.trim() === "") return;

    const newCat = {
      value: newCategory.toLowerCase().replace(/\s+/g, '-'),
      label: newCategory,
    };

    setCategories([...categories, newCat]);
    setTask({ ...task, category: newCat.value });
    handleCloseModal();
  };
  
    const handleChange = (
      e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>
    ) => {
      const { name, value } = e.target;
      setTask((prev) => ({ ...prev, [name]: value }));
    };
  
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
    console.log("Updated Task:", task);
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
    };
  
    const handleCancel = () => {
      navigate("/taskdashboard");
    };
  
    return (
      <div className="d-flex justify-content-center align-items-start w-100 h-100 pt-4">
      <CCard style={{ maxWidth: '600px', width: '100%' }}>
        <CCardHeader className="d-flex justify-content-between align-items-center">
          <strong>Update Task</strong>
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
  
        <CForm onSubmit={handleSubmit}>
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
                  required
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
                      <CIcon icon={cilPlus} />
                    </CButton>
                  </CTooltip>
                </CInputGroup>
              </CCol>
              <CCol md={6} className="mb-3">
                <CFormLabel htmlFor="dueDate">Due Date</CFormLabel>
                <CFormInput
                  type="date"
                  id="dueDate"
                  name="dueDate"
                  value={task.dueDate}
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
                  style={{ resize: 'none' }}
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
      </div>
    );
  }
  
export default UpdateTask;
  