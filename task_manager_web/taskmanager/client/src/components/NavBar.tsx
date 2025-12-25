interface Props {
  title: string;
  searchButton: string;
}

function NavBar({ title, searchButton }: Props) {
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light fixed-top shadow-sm">
      <div className="container-fluid">
        <span className="navbar-brand mb-0 h1">{title}</span>
        <div className="navbar-nav mx-auto d-flex flex-row gap-4">
          <a href="/" className="nav-link active" aria-current="page">
            Home
          </a>
          <a href="#" className="nav-link">
            About
          </a>
          <a href="#" className="nav-link">
            Contact
          </a>
          <a href="#" className="nav-link">
            Tasks
          </a>
        </div>
        <form role="search" className="d-flex">
          <input
            type="search"
            className="form-control me-2"
            placeholder="Search"
            aria-label="Search"
          />
          <button className="btn btn-outline-success" type="submit">
            {searchButton}
          </button>
        </form>
      </div>
    </nav>
  );
}

export default NavBar;
