import "@/styles/globals.css";
import { SessionProvider,useSession } from 'next-auth/react';
import Footer from '../components/Looding';
import Home from ".";
export default function App({ Component, pageProps }) {


  return (
    <SessionProvider session={pageProps.session}>
        {Component.auth ? (
        <Auth>
          <Component {...pageProps} />
        </Auth>
      ) : (
        <Component {...pageProps} />
      )}
         </SessionProvider>
  );
}
// AuthWrapper component to check session state
function Auth({ children }) {
  const { data: session, status } = useSession();
  
  if (status === "loading") {
    return <Loading/>; // 読み込み中の表示
  }

  if (!session) {
    // ログインしていない場合はログインページを表示
    return <Home />;
  }

  // ログインしている場合は元のコンテンツを表示
  return <>{children}</>;
}