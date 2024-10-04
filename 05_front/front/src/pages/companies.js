import { useState } from 'react';
import { useSession } from 'next-auth/react';
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });
export default function Companies() {
    const [query, setQuery] = useState('');
    const [results, setResults] = useState([]);
    const [registeredCompanies, setRegisteredCompanies] = useState([]); // 登録した会社のIdを管理
    const { data: session } = useSession();

    const searchCompanies = async () => {
        const response = await fetch(`http://localhost:1010/company?name=${query}`);
        console.log(response);
        const { companies } = await response.json();
        if (companies.length ==0){
        setResults(null);
        }else{

        setResults(companies);
        }
    };

    const addRecruit = async (company) => {
        // Extract JWT token from session
        console.log(session);
        const token = session?.jwt;  // Make sure the JWT token is stored in the session data

        if (!token) {
            console.error("JWT token not found");
            return;
        }

        const recruitData = {
            CompanyId: company.Id,  // 会社のIdを使用
            Name: company.Name,  // 会社名
            GoodPoint: '',  // ユーザーが入力する良い点
            SelfPr: '',  // ユーザーが入力する自己PR
            ConcernPoint: ''  // ユーザーが入力する懸念点
        };

        const response = await fetch(`http://localhost:1020/api/v1/recruits`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`  // JWTをAuthorizationヘッダーに追加
            },
            body: JSON.stringify(recruitData)  // recruitDataをJSON文字列に変換
        });

        const data = await response.json();  // Make sure to await response.json()
        if (response.ok) {
            // 登録成功の場合、登録した会社のIdを追加
            setRegisteredCompanies((prev) => [...prev, company.Id]); // company.Idを使って会社を特定
            console.log(data);
        } else {
            console.error("Failed to fetch companies:", data);
        }
    };

    return (
        <main
        className={`flex min-h-screen flex-col items-center justify-between p-24 ${inter.className}`}
      > 
<div className="bg-gray-100 flex flex-col items-center justify-center">
    <h1 className="text-3xl font-bold mb-6">会社検索</h1>
    <div className="mb-4">
        <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="会社名を入力"
            className="px-4 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
            onClick={searchCompanies}
            className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-md shadow-sm hover:bg-blue-600"
        >
            検索
        </button>
    </div>
    {
results == null && (<h2>検索結果がありません</h2>)
    }
    {results != null && results.length >0 && (
        <ul className="w-full mx-auto bg-white rounded-md shadow-md p-4">
            {results.map((company, index) => (
                <li key={index} className="border-b last:border-none py-4">
                    <div className="flex justify-between items-center">
                        <div>
                            <h2 className="text-lg font-semibold text-gray-900">{company.name}</h2>
                            <a
                                href={company.url}
                                className="text-sm text-blue-500 hover:underline"
                                target="_blank"
                                rel="noopener noreferrer"
                            >
                                {company.url}
                            </a>
                        </div>
                        <span
                            className={`inline-block px-2 py-1 text-xs font-semibold rounded ${company.active ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'}`}
                        >
                            {company.active ? 'Active' : 'Inactive'}
                            <button
                                onClick={() => addRecruit(company)}
                                className={`ml-2 px-4 py-2 rounded-md shadow-sm ${registeredCompanies.includes(company.Id)
                                    ? 'bg-green-500 text-white' // 登録済み
                                    : 'bg-blue-500 text-white hover:bg-blue-600' // 未登録
                                }`}
                            >
                                {registeredCompanies.includes(company.Id) ? '登録済み' : '登録'}
                            </button>
                        </span>
                    </div>
                </li>
            ))}
        </ul>
    )}
</div>

        </main>
    );
}

Companies.auth = true