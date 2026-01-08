import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import styles from "../components/SignIn.module.css";
//import { Link } from "react-router-dom";

function Register() {
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
            placeholder="Name"
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