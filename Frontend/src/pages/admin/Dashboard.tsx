import Navbar from "../../components/Navbar";

export default function DashboardAdmin() {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />

      <div className="p-8">
        {/* Mensaje de bienvenida */}
        <div className="bg-white shadow-md rounded-2xl p-6 mb-8 border-l-4 border-indigo-600">
          <h2 className="text-2xl font-semibold text-indigo-700 flex items-center gap-2">
            👋 Hola, <span className="text-gray-800">Admin</span>
          </h2>
          <p className="text-gray-600 mt-2">
            Bienvenido al panel de control. Aquí puedes gestionar créditos,
            usuarios y reportes del sistema.
          </p>
        </div>

        {/* Contenido del panel */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="bg-white rounded-2xl shadow p-6 hover:shadow-lg transition-all">
            <h3 className="font-semibold text-lg text-gray-800 mb-2">
              📊 Estadísticas
            </h3>
            <p className="text-gray-500 text-sm">
              Visualiza los datos más importantes sobre los créditos activos.
            </p>
          </div>

          <div className="bg-white rounded-2xl shadow p-6 hover:shadow-lg transition-all">
            <h3 className="font-semibold text-lg text-gray-800 mb-2">
              👥 Usuarios
            </h3>
            <p className="text-gray-500 text-sm">
              Administra operadores y permisos dentro del sistema.
            </p>
          </div>

          <div className="bg-white rounded-2xl shadow p-6 hover:shadow-lg transition-all">
            <h3 className="font-semibold text-lg text-gray-800 mb-2">
              💼 Créditos
            </h3>
            <p className="text-gray-500 text-sm">
              Revisa el estado de las solicitudes y aprobaciones de créditos.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
