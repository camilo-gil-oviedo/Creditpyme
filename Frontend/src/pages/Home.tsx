import { Link } from "react-router-dom";

export default function Home() {
  return (
    <div className="h-screen flex flex-col justify-center items-center bg-gray-50">
      <h1 className="text-4xl font-bold text-blue-600 mb-6">
        Sistema de Créditos Pyme
      </h1>
      <p className="text-gray-600 mb-8">Accede con tu rol para continuar</p>
      <Link
        to="/login"
        className="px-6 py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition"
      >
        Iniciar sesión
      </Link>
    </div>
  );
}
