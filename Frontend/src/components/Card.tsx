export default function Card({
  title,
  value,
}: {
  title: string;
  value: string;
}) {
  return (
    <div className="bg-white shadow-md rounded-xl p-6 text-center">
      <h3 className="text-lg font-semibold text-gray-600 mb-2">{title}</h3>
      <p className="text-3xl font-bold text-blue-600">{value}</p>
    </div>
  );
}
