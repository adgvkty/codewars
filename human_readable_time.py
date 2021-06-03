def format_duration(seconds):
    hours = 0
    minutes = 0
    days = 0
    res = []
    string = ''
    if seconds // 60:
        minutes = seconds // 60
        seconds -= minutes * 60
        seconds_time = str(seconds) + ' seconds'
        res.append([seconds, seconds_time])
    if minutes // 60:
        hours = minutes // 60
        minutes -= hours * 60
        minutes_time = str(minutes) + ' minutes'
        res.append([minutes, minutes_time])
    else:
        minutes_time = str(minutes) + ' minutes'
        res.append([minutes, minutes_time])
    if hours // 24:
        days = hours // 24
        hours -= days * 24
        hours_time = str(hours) + ' hours'
        days_time = str(days) + ' days'
        res.append([hours, hours_time])
        res.append([days, days_time])
    res = res[::-1]
    for el in res:
        print(el)
        if el[0] != 0:
            string = string + el[1] + ' '
    return string
    

print(format_duration(3244))