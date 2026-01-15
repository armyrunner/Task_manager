import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import styles from "../components/SignIn.module.css";
import { useAuth } from "../components/useAuth";
import { useNavigate } from "react-router-dom";
import { useState } from "react";



interface LoginCredentials {
  username: string;
  email: string;
  password: string;
}


function Register() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [formData, setFormData] = useState<LoginCredentials>({
    username: "",
    email: "",
    password: "",
  });
  const [loading,setloading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  
  const handleChange  = (e: React.ChangeEvent<HTMLInputElement) => {
    setFormData(prev => ({
      ...prev,
      [e.target.name]:e.target.value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (!formData.email || !formData.password){
      setError('Email or Password is invalid!!');
      return;
    }

    setloading(true);

    try{
      const credentials: LoginCredentials = { email, password};

      const resp = await fetch('http://localhost:8080/api/auth/login',{
        method: 'POST',
        headers: {'Content-Type':'application/json',},
        body: JSON.stringify(credentials),
        credentials: 'include'
      });

      const data = await resp.json();

      if(!resp.ok){
        throw new Error(data.message || data.Error || 'Login Failed');
      }
      
      if(data.access_token){
        localStorage.setItem('access_token',data.access_token)
        localStorage.setItem('refresh_token',data.refresh_token)
        login(data.user);
        navigate('/taskdashboard',{replace: true});
      }
      
    } catch(err){
      console.error('Login error:', err);
      setError(
        err instanceof Error ? err.message : 'Something went wrong. Please Try again!'
      );
    } finally{
      setloading(false);
    }
  }
  return (
    <div className={styles.pageWrapper}>
    <div
      className="content justify-content-center align-items-center d-flex shadow-lg"
      id="content"
    >
    <div className="col-md-6">
      <Form>
        <div className="header-text mb-4">
          <h1>Create Account</h1>
        </div>
        <div className="input-group mb-3">
          <input
            type="text"
            placeholder="username"
            className="form-control form-control-lg bg-light fs-6"
          ></input>
        </div>
        <div className="input-group mb-3">
          <input
            type="email"
            placeholder="Email"
            className="form-control form-control-lg bg-light fs-6"
          ></input>
        </div>
        <div className="input-group mb-3 justify-content-center">
          <input
            type="password"
            placeholder="Password"
            className="form-control form-control-lg bg-light fs-6"
          ></input>
        </div>
        <Button
          variant="success"
          type="submit"
          className="btn border-white text-white w-50 fs-6"
          onClick={handleSubmit}
        >
          Register
        </Button>
      </Form>
    </div>
    </div>
    </div>
  );
}


export default Register;