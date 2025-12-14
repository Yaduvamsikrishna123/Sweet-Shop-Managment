import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { Candy, LogOut, User, LayoutDashboard } from 'lucide-react';

const Navbar = () => {
    const { user, logout } = useAuth();

    return (
        <nav className="bg-white shadow-md sticky top-0 z-50">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between h-16">
                    <div className="flex items-center">
                        <Link to="/" className="flex items-center space-x-2">
                            <Candy className="h-8 w-8 text-pink-600" />
                            <span className="text-xl font-bold text-gray-800">Sweet Shop</span>
                        </Link>
                    </div>
                    <div className="flex items-center space-x-4">
                        {user ? (
                            <>
                                <span className="text-gray-600 flex items-center">
                                    <User className="h-4 w-4 mr-1" /> {user.username}
                                </span>
                                {user.role === 'admin' && (
                                    <Link to="/admin" className="text-gray-600 hover:text-pink-600 flex items-center">
                                        <LayoutDashboard className="h-5 w-5 mr-1" /> Admin
                                    </Link>
                                )}
                                <button
                                    onClick={logout}
                                    className="text-gray-600 hover:text-red-600 flex items-center"
                                >
                                    <LogOut className="h-5 w-5 mr-1" /> Logout
                                </button>
                            </>
                        ) : (
                            <>
                                <Link to="/login" className="text-gray-600 hover:text-pink-600 font-medium">Login</Link>
                                <Link to="/register" className="bg-pink-600 text-white px-4 py-2 rounded-md hover:bg-pink-700 transition-colors">Register</Link>
                            </>
                        )}
                    </div>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
