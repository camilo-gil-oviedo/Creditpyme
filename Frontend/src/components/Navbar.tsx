import { useNavigate } from "react-router-dom";
import { LogOut } from "lucide-react";

export default function Navbar() {
  const navigate = useNavigate();

  const handleLogout = () => {
    // Elimina el rol guardado
    localStorage.removeItem("role");

    // Redirige al login
    navigate("/login");
  };

  const role = localStorage.getItem("role");

  return (
    <nav className="flex items-center justify-between bg-white shadow-md px-6 py-4">
      {/* Logo / Título */}
      <div className="flex items-center gap-2">
        <img
          src="https://cdn-icons-png.flaticon.com/512/1041/1041916.png"
          alt="Logo"
          className="w-8 h-8"
        />
        <h1 className="text-xl font-bold text-indigo-700">Crédito Pyme</h1>
      </div>

      {/* Rol e ícono de salida */}
      <div className="flex items-center gap-4">
        <p className="text-gray-600 text-sm capitalize">
          Rol actual:{" "}
          <span className="font-semibold text-indigo-600">{role}</span>
        </p>

        <button
          onClick={handleLogout}
          className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-xl shadow-md transition-all duration-200 hover:scale-[1.02]"
        >
          <LogOut size={18} />
          <span>Cerrar sesión</span>
        </button>
      </div>
    </nav>
  );
}
