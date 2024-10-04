// pages/auth/signout.js
import { signOut } from 'next-auth/react';
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function SignOut() {
  const handleSignOut = async () => {
    await signOut({ redirect: true, callbackUrl: '/' });
  };

  return (
    <main className={`flex min-h-screen flex-col items-center justify-center p-6 bg-gray-100 ${inter.className}`}>
      <div className="w-full max-w-sm p-8 bg-white rounded-lg shadow-md">
        <h1 className="text-2xl font-bold mb-6 text-center">Sign Out</h1>
        <button onClick={handleSignOut} className="w-full py-2 px-4 bg-red-600 text-white font-semibold rounded-md shadow-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">Sign Out</button>
      </div>
    </main>
  );
}
