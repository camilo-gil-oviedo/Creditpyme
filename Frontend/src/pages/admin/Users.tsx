export default function Users() {
  const users = [
    { id: 1, nombre: "María López", rol: "Operador" },
    { id: 2, nombre: "Carlos Pérez", rol: "Administrador" },
  ];

  return (
    <div className="min-h-screen bg-gray-50 p-8 fade-in">
      <h1 className="text-3xl font-bold mb-6 text-gray-800">Usuarios</h1>
      <table className="table shadow-lg bg-white rounded-xl">
        <thead className="bg-indigo-100 text-left">
          <tr>
            <th className="p-3">ID</th>
            <th className="p-3">Nombre</th>
            <th className="p-3">Rol</th>
          </tr>
        </thead>
        <tbody>
          {users.map((u) => (
            <tr key={u.id} className="border-t hover:bg-indigo-50">
              <td className="p-3">{u.id}</td>
              <td className="p-3">{u.nombre}</td>
              <td className="p-3">{u.rol}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
