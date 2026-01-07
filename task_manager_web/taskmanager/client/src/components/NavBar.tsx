//import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "./useAuth";

interface Props {
  title: string;
}

function NavBar({ title}: Props) {
  const navigate = useNavigate();
  const { isLoggedIn, logout } = useAuth();

  const handleAuthClick = () => {
    if (isLoggedIn) {
      logout();
      navigate("/home");
    } else {
      navigate("/signin");
    }
  };

  return (
    <nav
      className="navbar navbar-expand-lg fixed-top shadow-sm"
      data-bs-theme="light"
    >
      <div className="container-fluid">
        {/* Brand */}
        <span className="navbar-brand mb-0 h1">{title}</span>

        {/* Toggler (needed for mobile) */}
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarContent"
          aria-controls="navbarContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>

        {/* Collapsible content – this is very important! */}
        <div className="collapse navbar-collapse" id="navbarContent">
          {/* Centered nav links */}
          <ul className="navbar-nav mx-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <Link to="/home" className="nav-link">
                Home
              </Link>
            </li>
            <li className="nav-item">
              <Link to="/about" className="nav-link">
                About Us
              </Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link">
                Contact
              </Link>
            </li>
            <li className="nav-item">
              <Link to="/overview" className="nav-link">
                Overview
              </Link>
            </li>
          </ul>

          {/* Right side button */}
          <div className="d-flex">
            <button
              onClick={handleAuthClick}
              className={`btn ${
                isLoggedIn ? "btn-outline-danger" : "btn-outline-success"
              }`}
              type="button"
            >
              {isLoggedIn ? "Sign Out" : "Sign In"}
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
}

export default NavBar;
