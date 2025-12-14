import { createContext, useContext, useState, useEffect } from 'react';
import toast from 'react-hot-toast';

const CartContext = createContext();

export const useCart = () => useContext(CartContext);

export const CartProvider = ({ children }) => {
    const [cart, setCart] = useState(() => {
        try {
            const savedCart = localStorage.getItem('cart');
            return savedCart ? JSON.parse(savedCart) : [];
        } catch (error) {
            console.error("Failed to parse cart from localStorage:", error);
            return [];
        }
    });

    useEffect(() => {
        localStorage.setItem('cart', JSON.stringify(cart));
    }, [cart]);

    const addToCart = (sweet) => {
        setCart(prevCart => {
            const existingItem = prevCart.find(item => item.ID === sweet.ID);
            if (existingItem) {
                if (existingItem.quantity >= sweet.quantity) {
                    toast.error('Cannot add more than available stock');
                    return prevCart;
                }
                toast.success('Updated quantity in cart');
                return prevCart.map(item =>
                    item.ID === sweet.ID
                        ? { ...item, quantity: item.quantity + 1 }
                        : item
                );
            }
            toast.success('Added to cart');
            return [...prevCart, { ...sweet, quantity: 1 }];
        });
    };

    const removeFromCart = (sweetId) => {
        setCart(prevCart => prevCart.filter(item => item.ID !== sweetId));
        toast.success('Removed from cart');
    };

    const updateQuantity = (sweetId, newQuantity, maxStock) => {
        if (newQuantity < 1) return;
        if (newQuantity > maxStock) {
            toast.error('Cannot exceed available stock');
            return;
        }
        setCart(prevCart =>
            prevCart.map(item =>
                item.ID === sweetId ? { ...item, quantity: newQuantity } : item
            )
        );
    };

    const clearCart = () => {
        setCart([]);
        localStorage.removeItem('cart');
    };

    const cartTotal = cart.reduce((total, item) => total + (item.price * item.quantity), 0);
    const cartCount = cart.reduce((count, item) => count + item.quantity, 0);

    return (
        <CartContext.Provider value={{
            cart,
            addToCart,
            removeFromCart,
            updateQuantity,
            clearCart,
            cartTotal,
            cartCount
        }}>
            {children}
        </CartContext.Provider>
    );
};
