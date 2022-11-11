import mongoose from "mongoose";

const AccountInfoSchema = new mongoose.Schema({
    name: String,
    surname: String,
    mail: String,
    telegramUsername: String
});

const ProfilePicSchema = new mongoose.Schema({
    image: { type: Buffer, required: false }   //tipo Buffer, è opzionale 
})

const FeaturesSchema = new mongoose.Schema({
    internetConnection: Boolean,
    bathroom: Boolean,
    heating: Boolean,
    airConditioner: Boolean,
    electricalOutlet: Boolean,
    tap: Boolean,
    bedLinens: Boolean,
    pillows: Boolean
})

const BedMutableInfoSchema = new mongoose.Schema({
    place: String,
    images: [Buffer],
    description: { type: String, required: false },
    features: FeaturesSchema,
    minimumDaysNotice: Number
})

const BedSchema = new mongoose.Schema({
    id: String,
    BedMutableInfo: BedMutableInfoSchema,
    dateAvailabes: [Date],         //0-90
    reviewCount: Number,
    averageEvaluation: Number      // min = 0, max = 50.
})

const BookingSchema = new mongoose.Schema({
    bed: BedSchema,
    date: Date
})

const ReviewSchema = new mongoose.Schema({
    evaluation: Number,            // min = 0, max = 50.
    comment: String
})

export class Models {
    public AccountInfo = mongoose.model('AccountInfo', AccountInfoSchema);
    public ProfileAPic = mongoose.model('ProfilePic', ProfilePicSchema);
    public Features = mongoose.model('Features', FeaturesSchema);
    public Booking = mongoose.model('Booking', BookingSchema);
    public BedMutableInfo = mongoose.model('BedMutableInfo', BedMutableInfoSchema);
    public Bed = mongoose.model('Bed', BedSchema);
    public Review = mongoose.model('Review', ReviewSchema);
    
}

