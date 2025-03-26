import { Routes, Route } from 'react-router-dom'
import Home from './pages/home/home.tsx'
import Auth from './pages/auth/auth.tsx'

function App() {

  return (
    <Routes>
        <Route path="/" element={<Home />}/>
        <Route path="/auth" element={<Auth/>}/>
    </Routes>
  )
}

export default App
