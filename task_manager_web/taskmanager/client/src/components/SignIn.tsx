import { useState} from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import { Link, useNavigate } from "react-router-dom";
import styles from "./SignIn.module.css";
import { useAuth } from "./useAuth";

interface LoginCredentials {
  email: string;
  password: string;
}


function SignIn() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading,setloading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  
  const handleRegisterClick = () => {
    navigate('/register');
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!email || !password){
      setError('Email or Password is invalid!!');
      return;
    }

    setloading(true);
    setError(null);

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
        <div className="col-md-6 d-flex justify-content-center">
          <Form onSubmit={handleSubmit}>
            <div className="header-text mb-4">
              <h1>Sign In</h1>
            </div>
            {error && (
              <div className="alert alert-danger" role="alert">
                {error}
              </div>
            )}
            <div className="input-group mb-3">
              <input
                type="email"
                placeholder="Email"
                className="form-control form-control-lg bg-light fs-6"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                disabled={loading}
              ></input>
            </div>
            <div className="input-group mb-3 justify-content-center">
              <input
                type="password"
                placeholder="Password"
                className="form-control form-control-lg bg-light fs-6"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                disabled={loading}
              ></input>
            </div>
            <div
              className="d-flex justify-content-between align-items-center mb-3"
              style={{ gap: ".50em" }}
            >
              <Form.Check
                type="checkbox"
                label="Remember Me!"
                className="text-secondary"
              />
              <Link to="/forgot" className="forgotpassordlink">
                <small>Forgot Password</small>
              </Link>
            </div>
            <div className="d-flex justify-content-center">
            <Button
              variant="success"
              type="submit"
              className="btn border-white text-white w-50 fs-6"
              style={{ marginRight: '10px' }}
              disabled={loading}
            >
              Log In
            </Button>
            <Button
              variant="primary"
              type="button"
              className="btn border-white text-white w-50 fs-6"
              onClick={handleRegisterClick}
              disabled={loading}
            >
              Register
            </Button>
            </div>
          </Form>
        </div>
      </div>
    </div>
  );
}

export default SignIn;
