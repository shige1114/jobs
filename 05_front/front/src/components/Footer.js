// components/Footer.js
const Footer = () => {
    return (
        <footer className="bg-gray-800 text-white p-4 mt-4">
            <div className="max-w-6xl mx-auto text-center">
                <p>© {new Date().getFullYear()} Jobs. </p>
            </div>
        </footer>
    );
};

export default Footer;
