import { useState } from 'react';
import { motion } from 'framer-motion';
import { ShoppingCart, Plus, Minus } from 'lucide-react';

const SweetCard = ({ sweet, onPurchase }) => {
    const [quantity, setQuantity] = useState(1);

    const handleIncrement = () => {
        if (quantity < sweet.quantity) {
            setQuantity(prev => prev + 1);
        }
    };

    const handleDecrement = () => {
        if (quantity > 1) {
            setQuantity(prev => prev - 1);
        }
    };

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

                <div className="flex items-center justify-between mb-4">
                    <span className={`text-sm font-medium ${sweet.quantity > 0 ? 'text-green-600' : 'text-red-600'}`}>
                        {sweet.quantity > 0 ? `${sweet.quantity} in stock` : 'Out of Stock'}
                    </span>

                    {sweet.quantity > 0 && (
                        <div className="flex items-center space-x-2 bg-gray-50 rounded-lg p-1">
                            <button
                                onClick={handleDecrement}
                                disabled={quantity <= 1}
                                className="p-1 hover:bg-white rounded-md transition-colors disabled:opacity-50"
                            >
                                <Minus className="h-4 w-4 text-gray-600" />
                            </button>
                            <span className="w-8 text-center font-medium">{quantity}</span>
                            <button
                                onClick={handleIncrement}
                                disabled={quantity >= sweet.quantity}
                                className="p-1 hover:bg-white rounded-md transition-colors disabled:opacity-50"
                            >
                                <Plus className="h-4 w-4 text-gray-600" />
                            </button>
                        </div>
                    )}
                </div>

                <button
                    onClick={() => onPurchase(quantity)}
                    disabled={sweet.quantity === 0}
                    className={`w-full flex items-center justify-center space-x-2 px-4 py-2 rounded-lg font-medium transition-colors ${sweet.quantity > 0
                        ? 'bg-pink-600 text-white hover:bg-pink-700'
                        : 'bg-gray-300 text-gray-500 cursor-not-allowed'
                        }`}
                >
                    <ShoppingCart className="h-4 w-4" />
                    <span>Add to Cart</span>
                </button>
            </div>
        </motion.div>
    );
};

export default SweetCard;
