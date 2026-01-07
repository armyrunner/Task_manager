import { BrowserRouter, Routes, Route } from 'react-router-dom';

import NavBar from './components/NavBar';

import SignIn from './components/SignIn';
import Register from './components/Register';
import HomePage from './components/HomePage';
import TaskDashboard from './components/TaskDashboard'
import { AuthProvider } from './components/AuthProvider';
import AboutPage from './components/AboutPage'
import ContactPage from './components/ContactPage'
import OverviewPage  from './components/OverviewPage'


//import PageNotFound from './components/PageNotFound'



import './App.css'

function App() {


  return (
    <>
    <AuthProvider>
    <BrowserRouter>
        <NavBar title='Task Manager' />
      <Routes>
        <Route path='/' element={<HomePage/>}/>
        <Route path='/home' element={<HomePage/>}/>
        <Route path='/signin' element={<SignIn/>}/>
        <Route path='/register' element={<Register/>}/>
        <Route path='/taskdashboard' element={<TaskDashboard/>}/>
        <Route path='/about' element={<AboutPage/>}/>
        <Route path='/contact' element={<ContactPage/>}/>
        <Route path='/overview' element={<OverviewPage/>}/>
      </Routes>
    </BrowserRouter>
    </AuthProvider>
    </>
  )
}

export default App
