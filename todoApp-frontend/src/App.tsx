import { Routes, Route } from 'react-router-dom'
import Home from './pages/home/home.tsx'
import Login from './pages/login/login.tsx'
import { CallBack } from './pages/login/callback.tsx'

function App() {

  return (
    
      <Routes>

        <Route path="/" element={<Home />}/>

        <Route path="/login" element={<Login />}/>
        
        <Route path="/login/callback" element={<CallBack />}/>

      </Routes>
    
  )
}

export default App
