import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import styles from "../components/SignIn.module.css";
import { useAuth } from "../components/useAuth";
import { useNavigate } from "react-router-dom";
import { useState } from "react";



interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}


function Register() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [formData, setFormData] = useState<RegisterRequest>({
    username: "",
    email: "",
    password: "",
  });
  const [loading,setloading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  
  const handleChange  = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData(prev => ({
      ...prev,
      [e.target.name]:e.target.value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (!formData.username || !formData.email || !formData.password){
      setError('All fields are required!!');
      return;
    }

    setloading(true);

    try{
      const registerCred: RegisterRequest = { username: formData.username, email: formData.email, password: formData.password};

      const resp = await fetch('http://localhost:8080/api/auth/register',{
        method: 'POST',
        headers: {'Content-Type':'application/json',},
        body: JSON.stringify(registerCred), 
        credentials: 'include'
      });

      const data = await resp.json();

      if(!resp.ok){
        throw new Error(data.message || data.Error || 'Registration Failed');
      }
      
      if(data.access_token){
        localStorage.setItem('access_token',data.access_token)
        localStorage.setItem('refresh_token',data.refresh_token)
        localStorage.setItem('user',JSON.stringify({id: data.user.id, username: data.user.username, email: data.user.email}));
        login({id: data.user.id, username: data.user.username, email: data.user.email});
        navigate('/taskdashboard',{replace: true});
      }
      
    } catch(err){
      console.error('Registration error:', err);
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
      <Form onSubmit={handleSubmit}>
        <div className="header-text mb-4">
          <h1>Create Account</h1>
        </div>
        {error && (
              <div className="alert alert-danger" role="alert">
                {error}
              </div>
            )}
        <div className="input-group mb-3">
          <input
            type="text"
            name="username"
            placeholder="username"
            className="form-control form-control-lg bg-light fs-6"
            value={formData.username}
            onChange={handleChange}
            disabled={loading}
            required
          ></input>
        </div>
        <div className="input-group mb-3">
          <input
            type="email"
            name="email"
            placeholder="Email"
            className="form-control form-control-lg bg-light fs-6"
            value={formData.email}
            onChange={handleChange}
            disabled={loading}
            required
          ></input>
        </div>
        <div className="input-group mb-3 justify-content-center">
          <input
            type="password"
            name="password"
            placeholder="Password"
            className="form-control form-control-lg bg-light fs-6"
            value={formData.password}
            onChange={handleChange}
            disabled={loading}
            required
          ></input>
        </div>
        <Button
          variant="success"
          type="submit"
          className="btn border-white text-white w-50 fs-6"
          disabled={loading}

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