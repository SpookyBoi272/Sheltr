import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import Home from "./pages/Home";
import Navbar from "./components/Navbar";
import Footer from "./components/Footer";
import Contacts from "./pages/Contacts";
import Landlords from "./pages/Landlords";
import Blog from "./pages/Blog";
import Register from "./components/Register";
import Fallback from "./pages/Fallback";

function App() {
  return (
    <div className="min-h-screen flex flex-col">
      <BrowserRouter>
        <Navbar />

          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/home" element={<Navigate to="/" />} />
            <Route path="/landlords" element={<Landlords />} />
            <Route path="/contacts" element={<Contacts />} />
            <Route path="/landlords" element={<Blog />} />
            <Route path="/landlords" element={<Blog />} />
            <Route path="/register" element={<Register />} />
            <Route path="*" element={<Fallback />} />
          </Routes>
        <Footer />
      </BrowserRouter>
    </div>
  );
}

export default App;
