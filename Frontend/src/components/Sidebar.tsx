import { motion } from "framer-motion";
import { Home, Users, FileText, Settings, LogOut } from "lucide-react";
import { useState } from "react";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const [active, setActive] = useState("Dashboard");

  const menu = [
    { name: "Dashboard", icon: <Home size={20} /> },
    { name: "Usuarios", icon: <Users size={20} /> },
    { name: "Solicitudes", icon: <FileText size={20} /> },
    { name: "Configuración", icon: <Settings size={20} /> },
  ];

  return (
    <div className="flex h-screen bg-gradient-to-br from-gray-50 to-gray-100">
      {/* Sidebar */}
      <aside className="w-64 bg-white/70 backdrop-blur-md shadow-md flex flex-col p-4 border-r border-gray-200">
        <div className="flex items-center gap-2 mb-8">
          <img src="/logo.png" alt="logo" className="w-8 h-8" />
          <h1 className="text-xl font-bold text-violet-700">Crédito Pyme</h1>
        </div>

        <nav className="space-y-2">
          {menu.map((item) => (
            <motion.button
              key={item.name}
              whileHover={{ scale: 1.05 }}
              onClick={() => setActive(item.name)}
              className={`flex items-center gap-3 w-full px-3 py-2 rounded-lg text-left transition-all duration-200
                ${
                  active === item.name
                    ? "bg-violet-100 text-violet-700 font-semibold"
                    : "text-gray-700 hover:bg-gray-100"
                }`}
            >
              {item.icon}
              {item.name}
            </motion.button>
          ))}
        </nav>
      </aside>

      {/* Main content */}
      <div className="flex-1 flex flex-col">
        {/* Header */}
        <header className="flex items-center justify-between bg-white/70 backdrop-blur-md shadow-sm px-6 py-4 border-b border-gray-200">
          <span className="text-gray-600">
            Rol actual:{" "}
            <strong className="text-violet-600">Administrador</strong>
          </span>
          <button className="flex items-center gap-2 bg-gradient-to-r from-violet-600 to-indigo-600 text-white px-4 py-2 rounded-xl shadow hover:shadow-md transition-all duration-300">
            <LogOut size={18} />
            Cerrar sesión
          </button>
        </header>

        {/* Content */}
        <main className="flex-1 overflow-y-auto p-8">
          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.4 }}
            className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3"
          >
            {children || (
              <>
                <Card title="Solicitudes Activas" value="128" color="violet" />
                <Card title="Usuarios Registrados" value="64" color="blue" />
                <Card title="Créditos Aprobados" value="82" color="emerald" />
              </>
            )}
          </motion.div>
        </main>
      </div>
    </div>
  );
}

function Card({
  title,
  value,
  color,
}: {
  title: string;
  value: string;
  color: string;
}) {
  const colors: Record<string, string> = {
    violet: "from-violet-500 to-indigo-500",
    blue: "from-sky-500 to-blue-500",
    emerald: "from-emerald-500 to-teal-500",
  };

  return (
    <motion.div
      whileHover={{ scale: 1.02 }}
      className="p-6 rounded-2xl bg-white/70 backdrop-blur-md shadow hover:shadow-lg transition-all duration-300"
    >
      <h2 className="text-gray-700 font-medium">{title}</h2>
      <p
        className={`text-4xl font-bold bg-gradient-to-r ${colors[color]} text-transparent bg-clip-text mt-2`}
      >
        {value}
      </p>
    </motion.div>
  );
}
