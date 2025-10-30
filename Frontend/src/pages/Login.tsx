import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";

export default function Login() {
  const navigate = useNavigate();
  const [user, setUser] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

 const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {

  e.preventDefault();
  setError(""); // limpiar errores previos

  try {
    // 🔹 Llamada al backend
    const response = await api.post("/login", {
      email: user,       // tu input "user" lo enviamos como email
      password: password, // contraseña
    });

    // Backend devuelve token y rol
    const { token, rol } = response.data;

    // Guardamos en localStorage
    localStorage.setItem("token", token);
    localStorage.setItem("role", rol);

    // Redirigir según rol
    if (rol === "admin") navigate("/admin");
    else if (rol === "operador") navigate("/operador");
    else navigate("/");

  } catch (err) {
    console.error("Login error:", err);
    setError("Credenciales incorrectas o servidor no disponible");
  }
};

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-indigo-600 via-purple-600 to-blue-600">
      <div className="bg-white bg-opacity-90 backdrop-blur-xl rounded-3xl shadow-2xl p-10 w-full max-w-md text-center">
        <h1 className="text-3xl font-bold text-indigo-700 mb-6">
          Crédito Pyme
        </h1>
        <p className="text-gray-500 mb-8">Inicia sesión para continuar</p>

        <form onSubmit={handleLogin} className="space-y-6">
          <input
            type="text"
            placeholder="Usuario"
            value={user}
            onChange={(e) => setUser(e.target.value)}
            className="w-full p-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:outline-none"
          />

          <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full p-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:outline-none"
          />

          {error && <p className="text-red-500 text-sm">{error}</p>}

          <button
            type="submit"
            className="btn w-full text-center font-semibold py-3 text-white bg-gradient-to-r from-indigo-600 to-indigo-800 rounded-xl hover:scale-[1.02] transition-transform"
          >
            Ingresar
          </button>
        </form>

        <div className="mt-6 text-sm text-gray-500">
          <p>
            <strong>Admin:</strong> usuario: <code>admin</code> / clave:{" "}
            <code>1234</code>
          </p>
          <p>
            <strong>Operador:</strong> usuario: <code>operador</code> / clave:{" "}
            <code>1234</code>
          </p>
        </div>
      </div>
    </div>
  );
}
