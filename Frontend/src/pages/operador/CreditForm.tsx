export default function CreditForm() {
  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <h1 className="text-3xl font-bold mb-6 text-gray-800">
        Formulario de Crédito
      </h1>

      <form className="bg-white p-6 rounded-xl shadow-md max-w-lg">
        <label className="block mb-2 text-gray-700">Nombre del cliente</label>
        <input
          type="text"
          className="w-full border border-gray-300 rounded-lg px-3 py-2 mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
        />

        <label className="block mb-2 text-gray-700">Monto solicitado</label>
        <input
          type="number"
          className="w-full border border-gray-300 rounded-lg px-3 py-2 mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
        />

        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700"
        >
          Enviar solicitud
        </button>
      </form>
    </div>
  );
}
