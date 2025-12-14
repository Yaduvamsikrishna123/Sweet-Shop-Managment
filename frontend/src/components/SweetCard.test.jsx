import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import SweetCard from '../components/SweetCard';

describe('SweetCard', () => {
    const mockSweet = {
        ID: 1,
        name: 'Gulab Jamun',
        category: 'Syrup Based',
        price: 15.0,
        quantity: 100,
    };

    const mockOnPurchase = vi.fn();

    it('renders sweet details correctly', () => {
        render(<SweetCard sweet={mockSweet} onPurchase={mockOnPurchase} />);

        expect(screen.getByText('Gulab Jamun')).toBeInTheDocument();
        expect(screen.getByText('Syrup Based')).toBeInTheDocument();
        expect(screen.getByText('$15.00')).toBeInTheDocument();
        expect(screen.getByText(/100 in stock/i)).toBeInTheDocument();
    });

    it('calls onPurchase with selected quantity when button is clicked', () => {
        render(<SweetCard sweet={mockSweet} onPurchase={mockOnPurchase} />);

        // Default quantity is 1
        const addButton = screen.getByText('Add to Cart');
        fireEvent.click(addButton);

        expect(mockOnPurchase).toHaveBeenCalledWith(1);
    });

    it('increments and decrements quantity correctly', () => {
        render(<SweetCard sweet={mockSweet} onPurchase={mockOnPurchase} />);

        const plusButton = screen.getByLabelText('Increase quantity');
        const minusButton = screen.getByLabelText('Decrease quantity');

        // Increment to 2
        fireEvent.click(plusButton);
        fireEvent.click(screen.getByText('Add to Cart'));
        expect(mockOnPurchase).toHaveBeenLastCalledWith(2);

        // Decrement back to 1
        fireEvent.click(minusButton);
        fireEvent.click(screen.getByText('Add to Cart'));
        expect(mockOnPurchase).toHaveBeenLastCalledWith(1);
    });

    it('does not decrement below 1', () => {
        render(<SweetCard sweet={mockSweet} onPurchase={mockOnPurchase} />);

        const minusButton = screen.getByLabelText('Decrease quantity');
        fireEvent.click(minusButton); // Should stay 1

        fireEvent.click(screen.getByText('Add to Cart'));
        expect(mockOnPurchase).toHaveBeenCalledWith(1);
    });

    it('does not increment above stock', () => {
        const lowStockSweet = { ...mockSweet, quantity: 2 };
        render(<SweetCard sweet={lowStockSweet} onPurchase={mockOnPurchase} />);

        const plusButton = screen.getByLabelText('Increase quantity');
        fireEvent.click(plusButton); // 2
        fireEvent.click(plusButton); // Should stay 2

        fireEvent.click(screen.getByText('Add to Cart'));
        expect(mockOnPurchase).toHaveBeenLastCalledWith(2);
    });
});
