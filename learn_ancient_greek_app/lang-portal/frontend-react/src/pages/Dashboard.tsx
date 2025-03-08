import { Link } from "react-router-dom";

const Dashboard = () => {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex flex-col md:flex-row items-center gap-8 mb-10">
        <div className="w-full md:w-1/3">
          <img 
            src="/images/high_priestess.jpg" 
            alt="Ancient Greek woman high priestess" 
            className="rounded-lg shadow-lg max-w-full h-auto"
          />
        </div>
        <div className="w-full md:w-2/3">
          <h1 className="text-3xl font-bold mb-4">Learn Ancient Greek with the High Priestess</h1>
          <p className="text-lg text-gray-700 mb-6">
            Embark on a journey to master the language of philosophers, poets, and historians. 
            Select an activity below to begin your learning adventure.
          </p>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {/* Translation Card */}
        <Link 
          to="/translate" 
          className="bg-white hover:bg-blue-50 transition-colors p-6 rounded-lg shadow border border-gray-200 text-center"
        >
          <h2 className="text-2xl font-semibold mb-2 text-blue-700">Translation</h2>
          <p className="text-gray-600 mb-4">Practice translating English to Ancient Greek</p>
          <div className="bg-blue-600 text-white py-2 px-4 rounded inline-block">
            Start Translating
          </div>
        </Link>

        {/* Listening Card */}
        <Link 
          to="/listening" 
          className="bg-white hover:bg-green-50 transition-colors p-6 rounded-lg shadow border border-gray-200 text-center"
        >
          <h2 className="text-2xl font-semibold mb-2 text-green-700">Listening</h2>
          <p className="text-gray-600 mb-4">Improve your Ancient Greek listening comprehension</p>
          <div className="bg-green-600 text-white py-2 px-4 rounded inline-block">
            Start Listening
          </div>
        </Link>

        {/* Writing Card */}
        <Link 
          to="/writing" 
          className="bg-white hover:bg-purple-50 transition-colors p-6 rounded-lg shadow border border-gray-200 text-center"
        >
          <h2 className="text-2xl font-semibold mb-2 text-purple-700">Writing</h2>
          <p className="text-gray-600 mb-4">Practice writing Ancient Greek with proper diacritics</p>
          <div className="bg-purple-600 text-white py-2 px-4 rounded inline-block">
            Start Writing
          </div>
        </Link>
      </div>
    </div>
  );
};

export default Dashboard;
