
// 1. Find High-Spending Users
db.orders.aggregate([
    {
        $group: {
            _id: "$userId",
            totalSpent: { $sum: "$totalAmount" }
        }
    },
    { $match: { totalSpent: { $gt: 500 } } },
    {
        $lookup: {
            from: "users",
            localField: "_id",
            foreignField: "userId",
            as: "userDetails"
        }
    }
]);

// 2. List Popular Products by Average Rating
db.products.aggregate([
    { $unwind: "$ratings" },
    {
        $group: {
            _id: "$productId",
            avgRating: { $avg: "$ratings.rating" }
        }
    },
    { $match: { avgRating: { $gte: 4 } } }
]);

// 3. Search for Orders in a Specific Time Range
db.orders.aggregate([
    { $match: { orderDate: { $gte: new Date("2024-12-01"), $lte: new Date("2024-12-31") } } },
    {
        $lookup: {
            from: "users",
            localField: "userId",
            foreignField: "userId",
            as: "userDetails"
        }
    }
]);

// 4. Update Stock After Order Completion
db.orders.find().forEach(order => {
    order.items.forEach(item => {
        db.products.updateOne(
            { productId: item.productId },
            { $inc: { stock: -item.quantity } }
        );
    });
});

// 5. Find Nearest Warehouse
db.warehouses.aggregate([
    {
        $geoNear: {
            near: { type: "Point", coordinates: [-74.006, 40.7128] },
            distanceField: "distance",
            maxDistance: 50000,
            query: { products: "P001" },
            spherical: true
        }
    }
]);
