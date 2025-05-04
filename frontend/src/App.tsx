import { BrowserRouter, Routes, Route } from "react-router-dom"
import Dashboard from "./pages/Dashboard"
import Register from "./pages/Register"
import Profile from "./pages/Profile"
import Login from "./pages/Login"
import NewChat from "./pages/NewChat"
import Chat from "./pages/Chat"
import './App.css'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={ <Dashboard/> } />
        <Route path="/chat/:id" element={ <Chat/> } />
        <Route path="/profile" element={ <Profile/> } />
        <Route path="/login" element={ <Login/> } />
        <Route path="/register" element={ <Register/> } />
        <Route path="/new-chat" element={ <NewChat/> } />
      </Routes>
    </BrowserRouter>
  )
}

export default App
