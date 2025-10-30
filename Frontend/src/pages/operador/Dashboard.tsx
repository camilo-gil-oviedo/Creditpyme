import Navbar from "../../components/Navbar";

export default function DashboardOperador() {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />

      <div className="p-8">
        {/* Mensaje de bienvenida */}
        <div className="bg-white shadow-md rounded-2xl p-6 mb-8 border-l-4 border-emerald-600">
          <h2 className="text-2xl font-semibold text-emerald-700 flex items-center gap-2">
            👷‍♂️ Hola, <span className="text-gray-800">Operador</span>
          </h2>
          <p className="text-gray-600 mt-2">
            ¡Listo para trabajar! Desde aquí puedes registrar solicitudes y
            actualizar información de clientes.
          </p>
        </div>

        {/* Contenido del panel */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="bg-white rounded-2xl shadow p-6 hover:shadow-lg transition-all">
            <h3 className="font-semibold text-lg text-gray-800 mb-2">
              📝 Solicitudes
            </h3>
            <p className="text-gray-500 text-sm">
              Crea, actualiza o revisa solicitudes de crédito en proceso.
            </p>
          </div>

          <div className="bg-white rounded-2xl shadow p-6 hover:shadow-lg transition-all">
            <h3 className="font-semibold text-lg text-gray-800 mb-2">
              📁 Clientes
            </h3>
            <p className="text-gray-500 text-sm">
              Consulta y edita la información de los clientes registrados.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
