import { motion } from 'framer-motion';
import { ShoppingCart } from 'lucide-react';

const SweetCard = ({ sweet, onPurchase }) => {
    return (
        <motion.div
            layout
            initial={{ opacity: 0, scale: 0.9 }}
            animate={{ opacity: 1, scale: 1 }}
            exit={{ opacity: 0, scale: 0.9 }}
            className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300"
        >
            <div className="h-48 bg-gradient-to-br from-pink-100 to-purple-100 flex items-center justify-center">
                <span className="text-6xl">🍬</span>
            </div>
            <div className="p-6">
                <div className="flex justify-between items-start mb-2">
                    <h3 className="text-xl font-bold text-gray-800">{sweet.name}</h3>
                    <span className="bg-pink-100 text-pink-800 text-xs px-2 py-1 rounded-full font-medium">
                        {sweet.category}
                    </span>
                </div>
                <p className="text-2xl font-bold text-pink-600 mb-4">${sweet.price.toFixed(2)}</p>
                <div className="flex justify-between items-center">
                    <span className={`text-sm font-medium ${sweet.quantity > 0 ? 'text-green-600' : 'text-red-600'}`}>
                        {sweet.quantity > 0 ? `${sweet.quantity} in stock` : 'Out of Stock'}
                    </span>
                    <button
                        onClick={() => onPurchase(sweet.ID)}
                        disabled={sweet.quantity === 0}
                        className={`flex items-center space-x-1 px-4 py-2 rounded-lg font-medium transition-colors ${sweet.quantity > 0
                                ? 'bg-pink-600 text-white hover:bg-pink-700'
                                : 'bg-gray-300 text-gray-500 cursor-not-allowed'
                            }`}
                    >
                        <ShoppingCart className="h-4 w-4" />
                        <span>Purchase</span>
                    </button>
                </div>
            </div>
        </motion.div>
    );
};

export default SweetCard;
