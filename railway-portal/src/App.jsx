import { useState } from 'react'
import { BrowserRouter, Route, Routes } from "react-router-dom";
import LandingPage from './pages/Landing';
import NotFound from './pages/NotFound';
import Booking from './pages/Booking';
import './App.css'
import Reservation from './pages/Reservation';


function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          flexGrow: 1,
        }}
      >
        <BrowserRouter>
          <Routes>
            <Route path="/" Component={LandingPage} />
            <Route path="/add-booking" Component={Booking} />
            <Route path="/view-bookings" Component={Reservation} />
            <Route path="*" Component={NotFound} />
          </Routes>
        </BrowserRouter>
      </div>
    </>
  )
}

export default App
