import Link from 'next/link';
import { useSession } from 'next-auth/react';

const Header = () => {
    const { data: session } = useSession(); // session を正しく取得

    return (
        <header className="bg-gray-800 text-white p-4">
            <div className="max-w-6xl mx-auto flex justify-between items-center">
                <h1 className="text-2xl font-bold">Jobs</h1>
                <nav>
                    <Link href="/" className="mr-4">Home</Link>

                    {session ? (
                        // セッションが存在する場合のリンク
                        <>
                            <Link href="/companies" className="mr-4">Company Search</Link>
                            <Link href="/recruits" className="mr-4">Recruits</Link>
                            <Link href="/auth/signout" className="mr-4">Sign Out</Link>
                        </>
                    ) : (
                        // セッションが存在しない場合のリンク
                        <Link href="/auth/signin" className="mr-4">Sign In</Link>
                    )}
                </nav>
            </div>
        </header>
    );
};

export default Header;
