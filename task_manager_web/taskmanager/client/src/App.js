
import Navbar from './component/navbar';
import Home from './component/home';
function App() {
  return (
    <div className="App">
      <Navbar/>
      <div className='content'>
        <Home/>
      </div>
    </div>
  );
}

export default App;
