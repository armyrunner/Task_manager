import { useState } from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import { Link, useNavigate } from "react-router-dom";
import styles from "./SignIn.module.css";
import { useAuth } from "./useAuth";


function SignIn() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  
  const handleRegisterClick = () => {
    navigate('/register');
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    
    const success = login(email, password);
    if (success) {
      navigate('/taskdashboard');
    } else {
      setError('Invalid email or password');
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
              ></input>
            </div>
            <div className="input-group mb-3 justify-content-center">
              <input
                type="password"
                placeholder="Password"
                className="form-control form-control-lg bg-light fs-6"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
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
            >
              Log In
            </Button>
            <Button
              variant="primary"
              type="button"
              className="btn border-white text-white w-50 fs-6"
              onClick={handleRegisterClick}
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
