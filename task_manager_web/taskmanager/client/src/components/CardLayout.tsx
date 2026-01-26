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
  CRow,
  CCol,
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
  const [error, setError] = useState<string | null>(null);
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchTasks = async () => {
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

    const data = await response.json();
    console.log("Fetched tasks:", data); // <--- check what you actually got

    if (!response.ok) {
      throw new Error(data.message || "Failed to fetch tasks");
    }

    // Ensure tasks is an array
    if (Array.isArray(data)) {
      setTasks(data);
    } else {
      console.warn("Tasks API did not return an array:", data);
      setTasks([]);
    }

  } catch (error) {
    console.error("Error fetching tasks:", error);
    setError("Failed to fetch tasks. Please try again.");
    setTasks([]);
  } finally {
    setLoading(false);
  }
};


  useEffect(() => {
    fetchTasks();
  }, []);

  const handleCardClick = (task: Task) => {
    setSelectedCard(task);
    setModalVisible(true);
  };

  const handleModalClose = () => {
    setModalVisible(false);
    setSelectedCard(null);
  };

  const handleUpdateTask = () => {
    navigate(`updatetask/${selectedCard?.id}`);
  };

  return (
    <div className={styles.cardGrid}>
      {loading && <CSpinner color="success" />}
      {error && <div className="alert alert-danger">Error: {error}</div>}

      {/* Responsive Grid */}
      <CRow xs={{ cols: 1 }} sm={{ cols: 2 }} md={{ cols: 3 }} lg={{ cols: 4 }} className="g-3">
        {tasks.map((task, index) => (
          <CCol key={task.id || index}>
            <CCard
              className={`${styles.cardMainContent} shadow h-100`}
              onClick={() => handleCardClick(task)}
            >
              <CCardBody className={styles.cardBody}>
                <CCardTitle className="border-bottom">{task.description}</CCardTitle>
                <CCardText>
                  <strong>Due Date:</strong> {task.due_date}<br />
                  <strong>Start Date:</strong> {task.start_date}<br />
                  <strong>Status:</strong> {task.status}<br />
                  <strong>Category:</strong> {task.category_name}<br />
                </CCardText>
              </CCardBody>
            </CCard>
          </CCol>
        ))}
      </CRow>

      {/* Modal */}
      <CModal visible={modalVisible} onClose={handleModalClose}>
        <CModalHeader>
          <CModalTitle>{selectedCard?.description}</CModalTitle>
        </CModalHeader>
        <CModalBody>
          <p><strong>Description:</strong> {selectedCard?.description}</p>
          <p><strong>Due Date:</strong> {selectedCard?.due_date}</p>
          <p><strong>Start Date:</strong> {selectedCard?.start_date}</p>
          <p><strong>Finish Date:</strong> {selectedCard?.finish_date}</p>
          <p><strong>Status:</strong> {selectedCard?.status}</p>
          <p><strong>Notes:</strong> {selectedCard?.notes}</p>
          <p><strong>Category:</strong> {selectedCard?.category_name}</p>
        </CModalBody>
        <CModalFooter>
          <CButton onClick={handleUpdateTask} color="primary" variant="outline">
            <CIcon icon={cilPenAlt} /> Update
          </CButton>
        </CModalFooter>
      </CModal>
    </div>
  );
}

export default CardLayout;
