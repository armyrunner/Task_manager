import { BrowserRouter, Routes, Route } from 'react-router-dom';

import NavBar from './components/NavBar';

import SigInRegister from './components/SignInRegister';
import HomePage from './components/HomePage';
//import AboutPage from './components/AboutPage'
//import ContactPage from './components/ContactPage'
//import OverViewPage  from './components/OverviewPage'
//import AboutPage from './components/AboutPage'

//import PageNotFound from './components/PageNotFound'



import './App.css'


function App() {


  return (
    <>
    <BrowserRouter>
        <NavBar title='Task Manager' signInButtonText='Sign In' />
      <Routes>
        <Route path='/home' element={<HomePage/>}/>
        <Route path='/signin' element={<SigInRegister/>}/>
        {/* <Route path='/about' element={<SignInPage/>}/> */}
        {/* <Route path='/about' element={<SignInPage/>}/> */}
        {/* <Route path='/about' element={<SignInPage/>}/> */}
      </Routes>
    </BrowserRouter>
    </>
  )
}

export default App
