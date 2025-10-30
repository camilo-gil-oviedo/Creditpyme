import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import AdminDashboard from "../pages/admin/Dashboard";
import OperadorDashboard from "../pages/operador/Dashboard";
import Users from "../pages/admin/Users";
import Sidebar from "../components/Sidebar";
import Navbar from "../components/Navbar";
import Login from "../pages/Login";

export default function AppRouter() {
  const role = localStorage.getItem("role");

  // Si no hay sesión activa, ir al login
  if (!role) {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="*" element={<Navigate to="/login" />} />
        </Routes>
      </BrowserRouter>
    );
  }

  // Layout principal (sidebar + navbar)
  return (
    <BrowserRouter>
      <div className="flex min-h-screen">
        <Sidebar children={undefined} />
        <div className="flex-1 flex flex-col">
          <Navbar />
          <Routes>
            {role === "admin" && (
              <>
                <Route path="/admin" element={<AdminDashboard />} />
                <Route path="/admin/users" element={<Users />} />
              </>
            )}
            {role === "operador" && (
              <Route path="/operador" element={<OperadorDashboard />} />
            )}
            <Route path="*" element={<Navigate to={`/${role}`} />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}
