import {
  CCard,
  CCardBody,
  CCardTitle,
  CCardText,
  CButton,
  CModal,
  CModalTitle,
  CModalFooter,
  CModalBody,
  CModalHeader,
  CSpinner,
} from "@coreui/react";
import styles from "./CardLayout.module.css";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { cilPenAlt } from "@coreui/icons";
import CIcon from "@coreui/icons-react";

interface Task {
  id: number;
  description: string;
  due_date: string;
  start_date: string;
  finish_date: string;
  status: string;
  notes: string;
  category_id: number;
  category_name: string;
}

function CardLayout() {
  const navigate = useNavigate();
  const [modalVisible, setModalVisible] = useState(false);
  const [selectedCard, setSelectedCard] = useState<Task | null>(null);
  const [error,setError] = useState<string | null>(null);
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading,setLoading] = useState<boolean>(false);

  const fetchTasks = async () => {
    setLoading(true);
    setError(null);
    try{
      const response = await fetch('http://localhost:8080/api/tasks',{
        method: 'GET',
        headers: {'Content-Type': 'application/json','Authorization': `Bearer ${localStorage.getItem('access_token')}`},
        credentials: 'include'
      });
      const data = await response.json();
      if(!response.ok){
        throw new Error(data.message || data.Error || 'Failed to fetch tasks');
      }
      setTasks(data);
    } catch (error) {
      console.error('Error fetching tasks:', error);
      setError('Failed to fetch tasks. Please try again.');
    } finally {
        setLoading(false);
    }

  }

  useEffect(() => {
    fetchTasks();
  },[]);

  const handleCardClick = (task: Task) => {
    setSelectedCard(task);
    setModalVisible(true);
  };

  const handleModalClose = () => {
    setModalVisible(false);
    setSelectedCard(null);
  };

  const handleUpdateTask = () => {
    navigate(`/updatetask/${selectedCard?.id}`);
  };

  return (
    <div className={styles.cardMainContent}>
      {loading && <CSpinner  color="success" />}
      {error && <div className="alert alert-danger">Error: {error}</div>}
      <div className="d-flex flex-wrap gap-3">
        {tasks.map((task) => (
          <CCard
            key={task.id}
            className={`${styles.cardMainContent} shadow`} 
            onClick={() => handleCardClick(task)}
          >
            <CCardBody className={styles.cardBody}>
              <CCardTitle className="border-bottom">{task.description}</CCardTitle>
              <CCardText>
                <strong>Due Date:</strong>{task.due_date}
                <strong>Start Date:</strong>{task.start_date}
                <strong>Status:</strong>{task.status}
                <strong>Category:</strong>{task.category_name}
                </CCardText>
             
            </CCardBody>
          </CCard>
        ))}
      </div>
      <CModal visible={modalVisible} onClose={handleModalClose}>
        <CModalHeader>
          <CModalTitle>{selectedCard?.description}</CModalTitle>
        </CModalHeader>
        <CModalBody>
          <p>
            <strong>Description:</strong> {selectedCard?.description}
          </p>
          <p>
            <strong>Due Date:</strong> {selectedCard?.due_date}
          </p>
          <p>
            <strong>Start Date:</strong> {selectedCard?.start_date}
          </p>
          <p>
            <strong>Finish Date:</strong> {selectedCard?.finish_date}
          </p>
          <p>
            <strong>Status:</strong> {selectedCard?.status}
          </p>
          <p>
            <strong>Notes:</strong> {selectedCard?.notes}
          </p>
          <p>
            <strong>Category:</strong> {selectedCard?.category_name}
          </p>
        </CModalBody>
        <CModalFooter>
          <CButton onClick={handleUpdateTask} color="primary" variant="outline">
            <CIcon icon={cilPenAlt} />
            Update
          </CButton>
        </CModalFooter>
      </CModal>
    </div>
  );
}

export default CardLayout;
