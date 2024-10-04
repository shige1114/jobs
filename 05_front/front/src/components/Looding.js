// pages/auth/signin.js
import { signIn } from 'next-auth/react';
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function Looding() {
  return (
    <main className={`flex min-h-screen flex-col items-center justify-center p-6 bg-gray-100 ${inter.className}`}>
        <h1 className="text-2xl font-bold mb-6 text-center">Loading</h1>
    </main>
  );
}
