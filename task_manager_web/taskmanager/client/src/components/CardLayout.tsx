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
} from "@coreui/react";
import styles from "./CardLayout.module.css";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { cilPenAlt } from "@coreui/icons";
import CIcon from "@coreui/icons-react";

interface Task {
  id: number;
  title: string;
  description: string;
  dueDate: string;
  startDate: string;
  finishDate: string;
  status: string;
  notes: string;
}

function CardLayout() {
  const navigate = useNavigate();
  const [modalVisible, setModalVisible] = useState(false);
  const [selectedCard, setSelectedCard] = useState<Task | null>(null);

  const tasks: Task[] = [
    {
      id: 1,
      title: "Task 1",
      description: "Task 1 description",
      dueDate: "2026-01-01",
      startDate: "2026-01-01",
      finishDate: "2026-01-01",
      status: "pending",
      notes: "Some notes for task 1",
    },
    {
      id: 2,
      title: "Task 2",
      description: "Task 2 description",
      dueDate: "2026-01-02",
      startDate: "2026-01-02",
      finishDate: "2026-01-02",
      status: "pending",
      notes: "Some notes for task 2",
    },
  ];

  const handleCardClick = (task: Task) => {
    setSelectedCard(task);
    setModalVisible(true);
  };

  const handleModalClose = () => {
    setModalVisible(false);
    setSelectedCard(null);
  };

  const handleUpdateTask = () => {
    navigate("updatetask");
  };

  return (
    <div className={styles.cardMainContent}>
      <div className="d-flex flex-wrap gap-3">
        {tasks.map((task) => (
          <CCard
            key={task.id}
            className={`${styles.cardMainContent} shadow`} 
            onClick={() => handleCardClick(task)}
          >
            <CCardBody className={styles.cardBody}>
              <CCardTitle className="border-bottom">{task.title}</CCardTitle>
              <p className={styles.cardBodyText}><strong>Description:</strong></p>
              <CCardText>{task.description}</CCardText>
             
            </CCardBody>
          </CCard>
        ))}
      </div>
      <CModal visible={modalVisible} onClose={handleModalClose}>
        <CModalHeader>
          <CModalTitle>{selectedCard?.title}</CModalTitle>
        </CModalHeader>
        <CModalBody>
          <p>
            <strong>Description:</strong> {selectedCard?.description}
          </p>
          <p>
            <strong>Due Date:</strong> {selectedCard?.dueDate}
          </p>
          <p>
            <strong>Start Date:</strong> {selectedCard?.startDate}
          </p>
          <p>
            <strong>Finish Date:</strong> {selectedCard?.finishDate}
          </p>
          <p>
            <strong>Status:</strong> {selectedCard?.status}
          </p>
          <p>
            <strong>Notes:</strong> {selectedCard?.notes}
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
