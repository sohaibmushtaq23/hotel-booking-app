export interface Client {
    id: number
    clientName: string
    cnic: string
    phone: string
    email: string
    discount: number
  }

export interface User {
    id: number
    userName: string
    password: string
    userRole: string
  }

export interface Room {
    id: number
    roomNo: string
    roomWidth: number
    roomLength: number
    doubleBeds: number
    singleBeds: number
    windows: number
    aC: boolean
    wifi: boolean
    hotWater: boolean
    balcony: boolean
    location: string
    roomCharges?: number
    roomImage?: string
    remarks?: string
    status?: string
}

export interface Booking{
    id: number
    idCustomer: number
    idRoom: number
    bookingStart: string|null
    bookingEnd: string|null
    extraCharges: number
    amountPaid: number
    reservedAt: string|null
    idReservedBy: number
    status: string
}
