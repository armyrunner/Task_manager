import {
    CTable,
    CCard,
    CCardBody,
    CCardHeader,
    CSpinner,
    CCardTitle,
  } from "@coreui/react";
  
  import { useEffect, useState } from "react";

  
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
  
  function CompletedTasks() {
    // Data state
    const [items, setItems] = useState<Task[]>([]);
  
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
    };
  
    const columns = Object.entries(columnConfig).map(([key, label]) => ({
      key,
      label,
      _props: { scope: "col" },
    }));
  
    const fetchCompletedTasks = async () => {
      setLoading(true);
      setError(null);
      try {
        const response = await fetch(
          "http://localhost:8080/api/tasks/completed",
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${localStorage.getItem("access_token")}`,
            },
          }
        );
        if (response.ok) {
          const data = await response.json();
          return data;
        }
      } catch (err) {
        console.error("Error fetching initial tasks:", err);
      }
    };

  
    useEffect(() => {
      fetchCompletedTasks().then((data) => {
        setItems(data || []);
        setLoading(false);
      });
    }, []);
  
    return (
      <div className="d-flex justify-content-center align-items-start w-100 h-100 pt-4">
        <CCard style={{ maxWidth: "100%", width: "100%" }}>
          <CCardHeader>
              <CCardTitle>Completed Task List</CCardTitle>
          </CCardHeader>
          <CCardBody>
            {loading && <CSpinner color="success" />}
            {error && <div className="alert alert-danger">Error: {error}</div>}
            <CTable
              columns={columns}
              items={items}
              tableHeadProps={{ color: "light" }}
              striped
              hover
              bordered
            />
          </CCardBody>
        </CCard>
      </div>
    );
  }
  
  export default CompletedTasks;
  