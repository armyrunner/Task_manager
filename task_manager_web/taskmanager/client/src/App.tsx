//import { useState } from 'react'

import './App.css'
import NavBar from './components/NavBar'
import Cover from './components/Cover'

function App() {


  return (
    <>
      <NavBar title='Task Manager' searchButton='Search'/>
      <Cover/>
    </>
  )
}

export default App
