//import { useState } from 'react'

import './App.css'
import NavBar from './components/NavBar'
import Cover from './components/Cover'

function App() {


  return (
    <>
      <NavBar title='Task Manager' signInButton='Sign In'/>
      <Cover/>
    </>
  )
}

export default App
