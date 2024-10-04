import { useEffect, useState } from 'react';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { Inter } from 'next/font/google'; // 例: Interフォントをインポート
const inter = Inter({ subsets: ['latin'] }); // フォントの設定

export default function Recruits() {
    const router = useRouter();
    const { tab = 'tab1' } = router.query;

    const { data: session } = useSession();
    const [recruits, setRecruits] = useState([]);
    const [selectedRecruit, setSelectedRecruit] = useState(null);
    const [formData, setFormData] = useState({
        GoodPoint: "",
        SelfPR: "",
        ConcernPoint: ""
    });

    useEffect(() => {
        const fetchRecruits = async () => {
            if (!session) return;
            const token = session?.jwt;

            const response = await fetch(`http://localhost:1020/api/v1/recruits`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
            });

            const { Recruits } = await response.json();
            setRecruits(Recruits);
            setSelectedRecruit(Recruits[0]); // 初期表示
            setFormData({
                GoodPoint: Recruits[0].GoodPoint,
                SelfPR: Recruits[0].SelfPR,
                ConcernPoint: Recruits[0].ConcernPoint,
            });
        };

        fetchRecruits();
    }, [session]);

    const handleSelectRecruit = (recruit) => {
        setSelectedRecruit(recruit);
        setFormData({
            GoodPoint: recruit.GoodPoint,
            SelfPR: recruit.SelfPR,
            ConcernPoint: recruit.ConcernPoint,
        });
    };

    const handleTabChange = (newTab) => {
        router.push({
            query: { ...router.query, tab: newTab },
        });
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value
        }));
    };

    const handleUpdate = async () => {
        if (!selectedRecruit) return;
        const token = session?.jwt;

        const response = await fetch(`http://localhost:1020/api/v1/recruits`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({
                Id: selectedRecruit.Id,
                GoodPoint: formData.GoodPoint,
                SelfPR: formData.SelfPR,
                ConcernPoint: formData.ConcernPoint,
            })
        });

        if (!response.ok) {
            alert('Update failed!');
        }else{
            console.log(response)
        }
    };

    return (
<main
    className={`flex min-h-screen flex-col items-center justify-between p-24 ${inter.className}`}
> 

    <div className="mx-auto p-6 bg-white rounded-lg shadow-md w-full">
    <h1 className="text-2xl font-bold text-gray-800 mb-4">就活管理</h1>
        {recruits.lenght == 0 ? <h2 className="text-2xl font-bold text-gray-800 mb-4">登録された会社はありません。</h2>:<></>}
        <div className="flex flex-wrap mb-4 border-b-2 border-gray-300 pb-2">
            {recruits.map((recruit, index) => (
                <button
                    key={index}
                    className={`w-32 px-4 py-2 text-sm font-semibold rounded-md ${selectedRecruit?.Id === recruit.Id ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-800 hover:bg-gray-300'} mr-2`}
                    onClick={() => {
                        handleSelectRecruit(recruit);
                        handleTabChange(`tab${index + 1}`);
                    }}
                >
                    {recruit.Name}
                </button>
            ))}
        </div>

        {selectedRecruit && (
            <div className="mt-4 p-4 border border-black rounded-lg bg-gray-50">
                <h2 className="text-xl font-bold">Details for {selectedRecruit.Name}</h2>

                <h3 className="text-lg font-semibold">Self PR:</h3>
                <textarea
                    name="SelfPR"
                    value={formData.SelfPR}
                    onChange={handleChange}
                    className="w-full h-24 p-2 border border-gray-300 rounded-md resize-none"
                />

                <h3 className="text-lg font-semibold">Good Points:</h3>
                <textarea
                    name="GoodPoint"
                    value={formData.GoodPoint}
                    onChange={handleChange}
                    className="w-full h-24 p-2 border border-gray-300 rounded-md resize-none"
                />

                <h3 className="text-lg font-semibold">Concern Points:</h3>
                <textarea
                    name="ConcernPoint"
                    value={formData.ConcernPoint}
                    onChange={handleChange}
                    className="w-full h-24 p-2 border border-gray-300 rounded-md resize-none"
                />

                <button
                    type="button"
                    onClick={handleUpdate}
                    className="mt-4 w-full px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
                >
                    Update
                </button>
            </div>
        )}
    </div>
</main>


    );
    
}

Recruits.auth = true