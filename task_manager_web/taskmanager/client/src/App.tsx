import { BrowserRouter, Routes, Route } from 'react-router-dom';

import NavBar from './components/NavBar';

import SignIn from './routes/SignIn';
import Register from './routes/Register';
import HomePage from './routes/HomePage';
import TaskDashboard from './components/TaskDashboard'
import { AuthProvider } from './components/AuthProvider';
import AboutPage from './routes/AboutPage'
import ContactPage from './routes/ContactPage'
import OverviewPage  from './routes/OverviewPage'
import AddTask from './routes/AddTask'
import UpdateTask from './routes/UpdateTask'
import DeleteTask from './routes/DeleteTask'
import CardLayout from './components/CardLayout'

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
        <Route path='/taskdashboard' element={<TaskDashboard/>}>
          <Route path='addtask' element={<AddTask/>}/>
          <Route path='updatetask' element={<UpdateTask/>}/>
          <Route path='deletetask' element={<DeleteTask/>}/>
          <Route index element={<CardLayout/>}/>
        </Route>
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
