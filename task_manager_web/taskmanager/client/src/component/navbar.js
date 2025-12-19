



const Navbar = () => {
    return ( 

        <nav className="navbar">
            <h1>Task Managaer Tool</h1>
                <div className="links">
                    <a href="/">Home</a>
                    <a href="/create">New Task</a>
                </div>
        </nav>

     );
}
 
export default Navbar;


// Inline styeling <a href='/create' style={{}}></a>
// style={{
//     color: "white",
//     backgroundColor: '#f1356d',
//     borderRadius: '8px'
// }}