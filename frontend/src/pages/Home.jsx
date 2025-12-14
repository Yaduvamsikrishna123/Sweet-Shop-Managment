import { useState, useEffect } from 'react';
import api from '../api/axios';
import SweetCard from '../components/SweetCard';
import { Search, Filter } from 'lucide-react';
import toast from 'react-hot-toast';
import { motion, AnimatePresence } from 'framer-motion';
import { useCart } from '../context/CartContext';

const Home = () => {
    const [sweets, setSweets] = useState([]);
    const [loading, setLoading] = useState(true);
    const [searchQuery, setSearchQuery] = useState('');
    const [selectedCategory, setSelectedCategory] = useState('All');
    const [categories, setCategories] = useState(['All']);
    const { addToCart } = useCart();

    useEffect(() => {
        fetchSweets();
    }, []);

    const fetchSweets = async () => {
        try {
            const response = await api.get('/sweets');
            setSweets(response.data);
            // Extract unique categories
            const uniqueCategories = ['All', ...new Set(response.data.map(s => s.category))];
            setCategories(uniqueCategories);
        } catch (error) {
            toast.error('Failed to fetch sweets');
        } finally {
            setLoading(false);
        }
    };

    const handleAddToCart = (sweet) => {
        addToCart(sweet);
    };

    const filteredSweets = sweets.filter(sweet => {
        const matchesSearch = sweet.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
            sweet.category.toLowerCase().includes(searchQuery.toLowerCase());
        const matchesCategory = selectedCategory === 'All' || sweet.category === selectedCategory;
        return matchesSearch && matchesCategory;
    });

    if (loading) {
        return (
            <div className="min-h-screen flex items-center justify-center">
                <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-pink-600"></div>
            </div>
        );
    }

    return (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            {/* Hero Section */}
            <div className="text-center mb-12">
                <h1 className="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
                    Welcome to the <span className="text-pink-600">Sweet Shop</span>
                </h1>
                <p className="text-lg text-gray-600 max-w-2xl mx-auto">
                    Discover our delicious collection of handcrafted sweets. From traditional favorites to modern delights.
                </p>
            </div>

            {/* Search and Filter */}
            <div className="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
                <div className="relative w-full md:w-96">
                    <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5" />
                    <input
                        type="text"
                        placeholder="Search sweets..."
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                        className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-transparent outline-none"
                    />
                </div>
                <div className="flex items-center space-x-2 w-full md:w-auto overflow-x-auto pb-2 md:pb-0">
                    <Filter className="text-gray-400 h-5 w-5 flex-shrink-0" />
                    {categories.map(category => (
                        <button
                            key={category}
                            onClick={() => setSelectedCategory(category)}
                            className={`px-4 py-2 rounded-full text-sm font-medium whitespace-nowrap transition-colors ${selectedCategory === category
                                ? 'bg-pink-600 text-white'
                                : 'bg-white text-gray-600 hover:bg-gray-100 border border-gray-200'
                                }`}
                        >
                            {category}
                        </button>
                    ))}
                </div>
            </div>

            {/* Sweets Grid */}
            <motion.div layout className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                <AnimatePresence>
                    {filteredSweets.map(sweet => (
                        <SweetCard key={sweet.ID} sweet={sweet} onPurchase={() => handleAddToCart(sweet)} />
                    ))}
                </AnimatePresence>
            </motion.div>

            {filteredSweets.length === 0 && (
                <div className="text-center py-12">
                    <p className="text-gray-500 text-lg">No sweets found matching your criteria.</p>
                </div>
            )}
        </div>
    );
};

export default Home;
