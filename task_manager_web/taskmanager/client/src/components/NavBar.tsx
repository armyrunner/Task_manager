interface Props {
    title: string;
    signInButton: string;
  }
  
  function NavBar({ title, signInButton }: Props) {
    return (
      <nav className="navbar navbar-expand-lg fixed-top shadow-sm" data-bs-theme="light">
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
                <a href="#home" className="nav-link">Home</a>
              </li>
              <li className="nav-item">
                <a href="#about" className="nav-link">About Us</a>
              </li>
              <li className="nav-item">
                <a href="#contact" className="nav-link">Contact</a>
              </li>
              <li className="nav-item">
                <a href="#overview" className="nav-link">Overview</a>
              </li>
            </ul>
  
            {/* Right side button */}
            <form role="search" className="d-flex">
              <button className="btn btn-outline-success" type="submit">
                {signInButton}
              </button>
            </form>
          </div>
        </div>
      </nav>
    );
  }
  
  export default NavBar;