import { redirect } from "next/navigation";

export default async function BidHub() {
  let message = "Under Construction!";

  try {
    const response = await fetch("http://localhost:6969/", {
      cache: "no-store",
    });
    if (response.ok) {
      message = await response.text();
    }
  } catch (error) {
    console.error("Failed to fetch from backend:", error);
  }

  return (
    <div className="flex items-center justify-center min-h-screen">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">BidHub</h1>
        <p className="text-xl text-gray-600">{message}</p>
      </div>
    </div>
  );
}
